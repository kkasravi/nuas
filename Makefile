PROJECT:=nuas
REPO:=kkasravi/$(PROJECT)
VERSION=v0.1.0
IMAGE:=$(REPO):$(VERSION)

.PHONY:  build run schemas

all: build

debug:
	@docker exec -e GODEBUGGER=$(GODEBUGGER) --privileged -it $(shell docker ps -q --filter ancestor=$(IMAGE)) /go/src/github.com/nervanasystems/nuas/scripts/godebug attach bin/apiserver

build:
	@echo building
	@docker build -t $(IMAGE) .

dev:
	@echo "Running image $(IMAGE) ..."
	@docker exec --privileged -it $(shell docker ps -q --filter ancestor=$(IMAGE)) /bin/bash

env-down:
	@echo "Stopping image $(shell docker ps -q --filter ancestor=$(IMAGE))"
	@docker stop $(shell docker ps -q --filter ancestor=$(IMAGE))

env-api:
	@echo "Starting image $(IMAGE) ..."
	@docker run --privileged -d -p 9443:9443 -t $(IMAGE) 

env-up:
	@echo "Bringing up apiserver $(IMAGE) ..."
	@docker run --privileged -d -p 9443:9443 -t $(IMAGE) /go/src/github.com/nervanasystems/nuas/scripts/run

logs:
	@echo "Fetch logs for $(IMAGE) ..."
	@docker logs -f $(shell docker ps -q --filter ancestor=$(IMAGE))

validate:
	@kubectl --kubeconfig=kubeconfig api-versions

create-jobtype:
	@kubectl --kubeconfig=kubeconfig create -f sample/jobtype.yaml

delete-jobtype:
	@kubectl --kubeconfig=kubeconfig delete -f sample/jobtype.yaml

clean:
	@echo "Removing image $(IMAGE) ..."
	@docker rmi $(IMAGE)

scrub:
	@echo "Removing <none> image ..."
	@docker images | grep '<none>' | awk '{print $$3}' | xargs docker rmi --force

schemas:
	@echo Generating schemas
	@kubectl proxy &
	@openapi2jsonschema http://127.0.0.1:8001/swagger.json

test:
	@echo Running tests
	@docker exec --privileged -it $(shell docker ps -q --filter ancestor=$(IMAGE)) go test ./pkg/...
