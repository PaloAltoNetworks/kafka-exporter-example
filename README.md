# Kafka Exporter

## Description

This repository contains the first steps of the Kafka exporter
It is based on the [events-exporter](https://github.com/aporeto-inc/events-exporter)

### How to develop

To adapt this code to whatever you will want to do, you need to

1. Clone the repository
2. Update the `main.go` file as you wish
3. Upload your docker image in a docker registry that is accessible to you (i.e: `DOCKER_REGISTRY=<docker.io/aporeto> PROJECT_NAME=<kafka-exporter> PROJECT_VERSION=<1.0.0> make push`)
4. Create the kubernetes and docker-swarm charts using the following command: `PROJECT_NAME=<kafka-exporter> PROJECT_VERSION=<1.0.0> make helm-repo`
5. Uncompress the tar file and serve it using `helm serve`
6. Install (i) or Upgrade (u) the chart from within your voila environment `deploy i/u <kafka-exporter>`

### Detailed instructions

Share you docker image

``` bash
DOCKER_REGISTRY=<docker.io/aporeto> PROJECT_NAME=<kafka-exporter> PROJECT_VERSION=<1.0.0> make push
```

Create the charts and aggregate them in a single `tar.gz` file:

``` bash
PROJECT_NAME=<kafka-exporter> PROJECT_VERSION=<1.0.0> make helm-repo
```

Uncompress the tar file to serve it from your local helm repo:

``` bash
tar xvf kafka-exporter-1.0.0-helm-local-repo.tgz
helm serve --repo-path=helm-local-repo --address 127.0.0.1:8880 &
helm repo add kafka-exporter http://127.0.0.1:8880
```

Activate your [voila](https://docs.console.aporeto.com/docs/install/what-is-voila/) environment:

``` bash
cd /path/to/your/voila/deployment
source conf.voila
```

#### Kubernetes

Install the app

``` bash
VOILA_HELM_REPO="kafka-exporter" deploy i kafka-exporter
```

and verify the pod is running:

``` bash
k get pods | grep kafka-exporter
```

#### Docker Swarm

install the app

``` bash
deploy i kafka-exporter/swarm-aporeto-kafka-exporter
```
