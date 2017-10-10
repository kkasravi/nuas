PROJECT:=nuas
REPO:=kkasravi/$(PROJECT)
VERSION=v0.1.0
IMAGE:=$(REPO):$(VERSION)

.PHONY:  build run schemas

all: dockerbuild

dockerdebug:
	@docker exec --privileged -it $(shell docker ps -q --filter ancestor=$(IMAGE)) -e GODEBUGGER=$(GODEBUGGER) /go/src/github.com/nervanasystems/nuas/scripts/godebug attach bin/apiserver

dockerbuild:
	@echo building
	@docker build -t $(IMAGE) .

dockershell:
	@echo "Running image $(IMAGE) ..."
	@docker exec --privileged -it $(shell docker ps -q --filter ancestor=$(IMAGE)) /bin/bash

dockerstop:
	@echo "Stopping image $(shell docker ps -q --filter ancestor=$(IMAGE))"
	@docker stop $(shell docker ps -q --filter ancestor=$(IMAGE))

dockerstart:
	@echo "Starting image $(IMAGE) ..."
	@docker run --privileged -d -p 9443:9443 -t $(IMAGE) 

dockerrun: dockerstart
	@echo "Starting apiserver in image $(IMAGE) ..."
	@docker exec --privileged -d $(shell docker ps -q --filter ancestor=$(IMAGE)) /go/src/github.com/nervanasystems/nuas/scripts/run

dockervalidate:
	@kubectl --kubeconfig=kubeconfig api-versions

dockerclean:
	@echo "Removing image $(IMAGE) ..."
	@docker rmi $(IMAGE)

dockerscrub:
	@echo "Removing <none> image ..."
	@docker images | grep '<none>' | awk '{print $$3}' | xargs docker rmi --force

schemas:
	@echo Generating schemas
	@kubectl proxy &
	@openapi2jsonschema http://127.0.0.1:8001/swagger.json

test:
	@echo Running tests
	@cd $(GOPATH)/src/github.com/nervanasystems/nuas && go test ./pkg/...
