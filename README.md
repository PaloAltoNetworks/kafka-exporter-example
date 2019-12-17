# Kafka Exporter

## Description

This application connects to Aporeto event system.
It collects all the events and log them on the standard output.

Important Notes:

1. Application should be running in the same environment as other Aporeto Services since NATS is not exposed publicly.
2. In order to have all the events reported in NATS, start your services with the parameter `EnableExportEvent`

This application can be enhanced to import all events whenever you feel it is necessary.
For instance, it could feed a Kafka server for real-time analysis.

## Running it

This application can be deployed using:

* as a kubernetes chart
* as a docker swarm chart

### Charts

Create the charts and aggregate them in a single `tar.gz` file:

``` bash
make helm-repo
```

Uncompress the tar file to serve it from your local helm repo:

``` bash
tar xvf kafka-exporter-1.0.0-helm-local-repo.tgz
helm repo add kafka-exporter http://127.0.0.1:8880
helm serve --repo-path=helm-local-repo --address 127.0.0.1:8880 &
```

> Note: Alternatively, you can publish the chart into an official/public helm repo.

## Voila Environment

Activate your [voila](https://docs.console.aporeto.com/docs/install/what-is-voila/) environment:

``` bash
cd /path/to/your/voila/deployment
source conf.voila
```

Then, tell the backend to publish all the events. By default, they do not expose the following information:

* flow reports
* file access reports
* audit reports
* event logs

In order to do that, you should set the appropriate environment variable:

``` bash
set_value global.integrations.externalEventsEnabled true
```

You can now restart the services and install the app beside your current Aporeto Deployment.

### Kubernetes

Restart `leon` and `zack`

``` bash
deploy u leon
deploy u zack
```

Install the app

``` bash
VOILA_HELM_REPO="kafka-exporter" deploy i kafka-exporter
```

### Docker Swarm

Restart `leon` and `zack`

``` bash
deploy u swarm-aporeto-leon
deploy u swarm-aporeto-zack
```

Install the app

``` bash
deploy i kafka-exporter/swarm-aporeto-kafka-exporter
```

### Publications

Publications are organized in different topics:

* `squall-events` contains all the requests the API is receiving. As this is more an internal topic, all services are using squall-events to discuss.
* `activities` contains a events that will be presented in the System Monitoring menu of the Aporeto UI.
* `external-auditreports` gather all reports of the enforcer's audit reports related to the system (syscalls).
* `external-fileaccessreports`contains all the file access reports reported by the enforcer.
* `external-flowreports` contains all the flows reported by the enforcer.
* `external-eventlogs` contains enforcer event logs.


To give you the context of a publication, here are the fields you need to understand:

* `identity` defines the type of the object related to this publication (enforcer, processingunit, externalnetwork ...)
* `operation` the operation that triggered the publication (create, update, retrieve, retrieve-many)`

See detailed [examples](EXAMPLES.md)

### Using Aporeto objects (Gaia)

All Aporeto objects are coming from the library named `gaia`. They all inherit a common interface named `elemental.Identifiable`.
That is why, any object has the property `identity` which contains the name of the Identity.

When receiving a publication, you can decode the publication into the corresponding Aporeto Object:

```go
// publicationToString converts a publication to a readable string
func publicationToIdentifiable(pub *bahamut.Publication) (elemental.Identifiable, error) {

  // Decode the event of the publication
  event := &elemental.Event{}

  if err := pub.Decode(event); err != nil {
    return "", fmt.Errorf("unable to decode the event %s", err)
  }


  // Decode the identifiable
  identifiable := gaia.Manager().IdentifiableFromString(event.Identity)
  if err := event.Decode(&identifiable); err != nil {
    return "", fmt.Errorf("unable to decode the identifiable %s", err)
  }

  return identifiable, nil
```

Later, you can use that `elemental.Identifiable` to do anything specific:

```go
  obj, err := publicationToIdentifiable(pub)
  if err != nil {
    return err
  }

  switch obj.Identity() {

  case gaia.ActivityIdentity:
    activity := obj.(*gaia.Activity)
    return sendActivityToMonitoringSystem(activity)

  case gaia.EnforcerIdentity:
    enforcer := obj.(*gaia.Enforcer)
    return doSomethingWithEnforcer(enfrocer)

  default:
    return fmt.Errorf("unsupported entity %s", gaia.ActivityIdentity.Name)

}
```
