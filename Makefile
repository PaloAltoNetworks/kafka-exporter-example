MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash -o pipefail

DOCKER_REGISTRY ?= docker.io/aporeto
PROJECT_NAME ?= kafka-exporter
PROJECT_SHA ?= $(shell git rev-parse HEAD)
PROJECT_VERSION ?= v0.0.0-dev
PROJECT_RELEASE ?= dev

init:
	dep ensure
	dep status

lint:
	@ go get -u github.com/golangci/golangci-lint/cmd/golangci-lint

	# this linter is disabled because of https://github.com/mvdan/unparam/issues/35
	# --enable=unparam
	golangci-lint run \
		--deadline=3m \
		--disable-all \
		--exclude-use-default=false \
		--enable=errcheck \
		--enable=goimports \
		--enable=ineffassign \
		--enable=govet \
		--enable=golint \
		--enable=unused \
		--enable=structcheck \
		--enable=varcheck \
		--enable=deadcode \
		--enable=unconvert \
		--enable=goconst \
		--enable=gosimple \
		--enable=misspell \
		--enable=staticcheck \
		--enable=prealloc \
		--enable=nakedret \
		--enable=typecheck \
		./...

test:
	go test ./... -race

build:
	go build

build_linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

package: build_linux
	cp ./kafka-exporter-example ./docker/app/kafka-exporter

container: package
	cd docker && docker build -t $(DOCKER_REGISTRY)/$(PROJECT_NAME):$(PROJECT_VERSION) .

push: container
	docker push $(DOCKER_REGISTRY)/$(PROJECT_NAME):$(PROJECT_VERSION)

charts: charts-k8s

charts-k8s:
	@rm -rf ./helm/repo
	@mkdir -p ./helm/repo
	@helm lint ./helm/charts/kafka-exporter -f ./helm/tests/values.yaml
	@helm package \
		--version $(PROJECT_VERSION) \
		--app-version $(PROJECT_VERSION) \
		./helm/charts/kafka-exporter \
		-d ./helm/repo/
		
helm-repo: charts
	@mkdir -p helm-local-repo
	@cp ./helm/repo/* ./swarm/repo/* helm-local-repo/
	@helm repo index helm-local-repo/
	@tar czf $(PROJECT_NAME)-$(PROJECT_VERSION)-helm-local-repo.tgz helm-local-repo
	@rm -rf helm-local-repo
