PROJECT:=nuas
REPO:=kkasravi/$(PROJECT)
VERSION=latest
IMAGE:=$(REPO):$(VERSION)

.PHONY:  build run schemas

all: build

debug:
	@docker-compose exec --privileged nuas env GODEBUGGER=$(GODEBUGGER) /go/src/github.com/nervanasystems/nuas/scripts/godebug attach bin/apiserver

build:
	@echo building
	@docker build -t $(IMAGE) .

dev:
	@echo "create shell in $(IMAGE) ..."
	@docker-compose exec --privileged nuas /bin/bash

env-down:
	@echo "Stopping nuas"
	@docker-compose down

env-api:
	@echo "Starting image $(IMAGE) ..."
	@docker-compose run nuas /bin/bash

env-up:
	@echo "Bringing up apiserver $(IMAGE) ..."
	@docker-compose up -d
	@docker-compose ps

logs:
	@echo "Fetch logs for $(IMAGE) ..."
	@docker-compose logs nuas

validate:
	@kubectl --kubeconfig=kubeconfig api-versions

create-jobtype:
	@docker-compose exec nuas kubectl --kubeconfig=kubeconfig create -f sample/jobtype.yaml

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
	@docker-compose exec nuas go test ./pkg/...
