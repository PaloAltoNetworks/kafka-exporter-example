# Kafka Exporter

## Description

This repository contains the first steps of the Kafka exporter

## How to build

To adapt this code to your liking then

```console
DOCKER_REGISTRY=<your.registry/your_organisation> PROJECT_VERSION=1.0.0 make
```

This will:

- build and push a docker images as `your.registry/your_organisation/kafka-exporter:2.0.0`
- build the Helm charts locally

## How to deploy

First you need to enable the push of external events to Nats on the targetted setup.

Activate your [voila](https://docs.console.aporeto.com/docs/install/what-is-voila/) environment then

```console
# Enable the external events
set_value global.externalEventsEnabled true override
# Apply the configuration
doit
```

Then push the Helm charts to your Helm repository then:

```console
# add you repository
helm repo add kafka-exporter <your_helm_repository_url>
# deploy
VOILA_HELM_REPO=kafka-exporter deploy i kafka-exporter
```

To update:

```console
# add you repository
helm repo add kafka-exporter <your_helm_repository_url>
# update
VOILA_HELM_REPO=kafka-exporter deploy u kafka-exporter
```

To uninstall just type `deploy d kafka-exporter`

> Note: you must deploy the kafka-exporter using the voila environement as explained above because
> it requires access to some secrets to be able to connect to Nats.
> If the docker image of the kafka-exporter is not hosted on the same registry where the contron plane
> docker images are hosted, add an option as below during install or update:
>
> ```console
> # use another docker registry
> VOILA_HELM_REPO=kafka-exporter deploy i kafka-exporter --set global.imageRegistry=<your.registry/your_organisation>
> ```
