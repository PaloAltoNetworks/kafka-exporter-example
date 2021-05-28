package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"go.aporeto.io/addedeffect/logutils"
	"go.aporeto.io/bahamut"
	"go.aporeto.io/elemental"
	"go.aporeto.io/gaia"
	"go.aporeto.io/kafka-exporter-example/internal/configuration"
	"go.uber.org/zap"
)

func main() {

	cfg := configuration.NewConfiguration()
	logutils.Configure(cfg.LogLevel, cfg.LogFormat)

	tlsConfig := &tls.Config{
		RootCAs:      cfg.CAPool,
		Certificates: cfg.ClientCertificates,
	}

	pubsub := bahamut.NewNATSPubSubClient(
		cfg.NATSAddress,
		bahamut.NATSOptClientID(cfg.NATSClientID),
		bahamut.NATSOptClusterID(cfg.NATSClusterID),
		bahamut.NATSOptCredentials(cfg.NATSUser, cfg.NATSPassword),
		bahamut.NATSOptTLS(tlsConfig),
	)
	defer pubsub.Disconnect() // nolint

	connectCtx, connectCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer connectCancel()

	if err := pubsub.Connect(connectCtx); err != nil {
		zap.L().Fatal("Could not connect to nats", zap.Error(err))
	}
	defer pubsub.Disconnect() // nolint

	zap.L().Info("Connected to nats", zap.String("address", cfg.NATSAddress))

	// Squall events
	eventsErrs := make(chan error)
	eventsPubs := make(chan *bahamut.Publication)
	pubsub.Subscribe(eventsPubs, eventsErrs, cfg.TopicEvents, bahamut.NATSOptSubscribeQueue(cfg.NATSQueueName))

	// Activities also named Audit logs
	activitiesErrs := make(chan error)
	activitiesPubs := make(chan *bahamut.Publication)
	pubsub.Subscribe(activitiesPubs, activitiesErrs, cfg.TopicActivities, bahamut.NATSOptSubscribeQueue(cfg.NATSQueueName))

	// Flow reports
	flowReportsErrs := make(chan error)
	flowReportsPubs := make(chan *bahamut.Publication)
	pubsub.Subscribe(flowReportsPubs, flowReportsErrs, cfg.TopicFlowReports, bahamut.NATSOptSubscribeQueue(cfg.NATSQueueName))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		select {

		// Squall events
		case p := <-eventsPubs:
			if err := sendPublication(p); err != nil {
				fmt.Println("sendPublication error: ", err)
			}
		case err := <-eventsErrs:
			fmt.Printf("Got an event error %s\n", err)

		// Activities
		case p := <-activitiesPubs:
			if err := sendActivity(p); err != nil {
				fmt.Println("sendActivity error: ", err)
			}
		case err := <-activitiesErrs:
			fmt.Printf("Got an activity error %s\n", err)

		// Flow reports
		case p := <-flowReportsPubs:
			if err := sendFlowReport(p); err != nil {
				fmt.Println("sendFlowReport error: ", err)
			}
		case err := <-flowReportsErrs:
			fmt.Printf("Got an flow record error %s\n", err)

		case <-ctx.Done():
			fmt.Println("Exit")
		}
	}
}

// sendPublication sends a publication based on an event.
func sendPublication(pub *bahamut.Publication) error {

	// Convert the publication to an Aporeto object
	obj, err := publicationToIdentifiable(pub)
	if err != nil {
		return err
	}

	switch obj.Identity() {

	// For now, we are only listening to enforcer events
	case gaia.EnforcerIdentity:
		enforcer := obj.(*gaia.Enforcer)
		return sendEnforcerReport(enforcer)

	default:
		fmt.Printf("skipping event for %s\n", obj.Identity().Name)
		return nil
	}
}

// publicationToString converts a publication to an identifiable
func publicationToIdentifiable(pub *bahamut.Publication) (elemental.Identifiable, error) {

	// Decode the event of the publication
	event := &elemental.Event{}

	if err := pub.Decode(event); err != nil {
		return nil, fmt.Errorf("unable to decode the publication: %s", err)
	}

	// Decode the identifiable
	identifiable := gaia.Manager().IdentifiableFromString(event.Identity)
	if err := event.Decode(&identifiable); err != nil {
		return nil, fmt.Errorf("unable to decode the event: %s", err)
	}

	return identifiable, nil
}

// sendEnforcerReport sends a report to Kafka
func sendEnforcerReport(enforcer *gaia.Enforcer) error {

	fmt.Printf(`--- Enforcer to be send to Kafka:
	ID: %s
	Name: %s
	Namespace: %s
	Operational Status: %s
	Enforcement Status: %s
`, enforcer.ID, enforcer.Name, enforcer.Namespace, enforcer.OperationalStatus, enforcer.EnforcementStatus)

	return nil
}

// sendActivity sends an activity from a publication.
func sendActivity(pub *bahamut.Publication) error {

	activity := gaia.NewActivity()

	if err := pub.Decode(activity); err != nil {
		return fmt.Errorf("unable to decode the activity: %s", err)
	}

	if activity.TargetIdentity != gaia.EnforcerIdentity.Name {
		return fmt.Errorf("skipping activity for target identity: %s", activity.TargetIdentity)
	}

	fmt.Printf(`--- Activity to be send to Kafka:
	ID: %s
	Date: %s
	Namespace: %s
	Diff: %s
`, activity.ID, activity.Date, activity.Namespace, activity.Diff)

	return nil
}

// sendFlowReport sends a flow report from a publication.
func sendFlowReport(pub *bahamut.Publication) error {

	flowReport := gaia.NewFlowReport()

	if err := pub.Decode(flowReport); err != nil {
		return fmt.Errorf("unable to decode the flowReport: %s", err)
	}

	fmt.Printf(`--- FlowReport to be send to Kafka:
	Date: %v
	Namespace: %s
	SourceID: %s
	DestinationID: %s
	Hits: %d
`, flowReport.Timestamp, flowReport.Namespace, flowReport.SourceID, flowReport.DestinationID, flowReport.Value)

	return nil
}
