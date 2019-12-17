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

	if connected := pubsub.Connect().Wait(30 * time.Second); !connected {
		zap.L().Fatal("Could not connect to nats")
	}

	zap.L().Info("Connected to nats", zap.String("address", cfg.NATSAddress))

	// Squall events
	eventsErrs := make(chan error)
	eventsPubs := make(chan *bahamut.Publication)
	pubsub.Subscribe(eventsPubs, eventsErrs, "squall-events", bahamut.NATSOptSubscribeQueue(cfg.NATSQueueName))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for {
		select {

		case p := <-eventsPubs:
			if errSend := sendPublication(p); errSend != nil {
				fmt.Println("sendPublication error: ", errSend)
			}

		case err := <-eventsErrs:
			fmt.Printf("Got an event error %s\n", err)

		case <-ctx.Done():
			fmt.Println("Exit")
		}
	}
}

// sendPublication sends a publication
func sendPublication(pub *bahamut.Publication) error {

	obj, err := publicationToIdentifiable(pub)
	if err != nil {
		return err
	}

	switch obj.Identity() {

	case gaia.EnforcerIdentity:
		enforcer := obj.(*gaia.Enforcer)
		return sendEnforcerReport(enforcer)

	default:
		return fmt.Errorf("unsupported entity %s", gaia.ActivityIdentity.Name)

	}
}

// publicationToString converts a publication to a readable string
func publicationToIdentifiable(pub *bahamut.Publication) (elemental.Identifiable, error) {

	// Decode the event of the publication
	event := &elemental.Event{}

	if err := pub.Decode(event); err != nil {
		return nil, fmt.Errorf("unable to decode the event %s", err)
	}

	// Decode the identifiable
	identifiable := gaia.Manager().IdentifiableFromString(event.Identity)
	if err := event.Decode(&identifiable); err != nil {
		return nil, fmt.Errorf("unable to decode the identifiable %s", err)
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
