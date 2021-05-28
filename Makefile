MAKEFLAGS += --warn-undefined-variables
SHELL := /bin/bash -o pipefail

DOCKER_REGISTRY ?= docker.io/aporeto
PROJECT_NAME ?= kafka-exporter
PROJECT_SHA ?= $(shell git rev-parse HEAD)
PROJECT_VERSION ?= v0.0.0-dev
PROJECT_RELEASE ?= dev

export GO111MODULE = on

define VERSIONS_FILE
package versions

// Various version information.
var (
	ProjectVersion = "$(PROJECT_VERSION)"
	ProjectSha     = "$(PROJECT_SHA)"
	ProjectRelease = "$(PROJECT_RELEASE)"
)
endef
export VERSIONS_FILE

default: push charts

version:
	@echo "$$VERSIONS_FILE" > ./internal/versions/versions.go

lint: version
	# linting
	golangci-lint run \
		--disable-all \
		--exclude-use-default=false \
		--enable=errcheck \
		--enable=goimports \
		--enable=ineffassign \
		--enable=revive \
		--enable=unused \
		--enable=structcheck \
		--enable=staticcheck \
		--enable=varcheck \
		--enable=deadcode \
		--enable=unconvert \
		--enable=misspell \
		--enable=prealloc \
		--enable=nakedret \
		--enable=typecheck \
		./...

test: lint
	# testing
	go test ./... -race

build: test
	# building
	go build

build_linux: test
	# building
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

container: build_linux
	# creating container
	@mkdir -p ./docker/app
	@cp ./kafka-exporter-example ./docker/app/kafka-exporter
	@cd docker && docker build -t $(DOCKER_REGISTRY)/$(PROJECT_NAME):$(PROJECT_VERSION) .

push: container
	# pushing container
	docker push $(DOCKER_REGISTRY)/$(PROJECT_NAME):$(PROJECT_VERSION)

charts:
	# building helm charts
	@rm -rf ./helm/repo
	@mkdir -p ./helm/repo
	@helm lint ./helm/charts/kafka-exporter -f ./helm/tests/values.yaml
	@helm package \
		--version $(PROJECT_VERSION) \
		--app-version $(PROJECT_VERSION) \
		./helm/charts/kafka-exporter \
		-d ./helm/repo/
