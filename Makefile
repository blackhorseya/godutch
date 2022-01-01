APP_NAME=godutch
VERSION=latest
PROJECT_ID=sean-side
NS=side
DEPLOY_TO=uat
REGISTRY=gcr.io
IMAGE_NAME=$(REGISTRY)/$(PROJECT_ID)/$(APP_NAME)
HELM_REPO_NAME=blackhorseya
CHART_NAME=godutch

DB_URI = "mysql://godutch:changeme@tcp(localhost:3308)/godutch?charset=utf8mb4&parseTime=True&loc=Local"

check_defined = $(if $(value $1),,$(error Undefined $1))

.PHONY: help # Generate list of targets with descriptions
help:
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/\1 \2/' | expand -t20

.PHONY: clean # remove data
clean:
	@rm -rf bin coverage.txt profile.out

.PHONY: lint # execute golint
lint:
	@golint ./...

.PHONY: report # refresh goreportcard
report:
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/godutch'

.PHONY: test-unit # execute unit test
test-unit:
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: build-image # build docker image with APP_NAME and VERSION
build-image:
	$(call check_defined,VERSION)
	@docker build -t $(IMAGE_NAME):$(VERSION) \
	--label "app.name=$(APP_NAME)" \
	--label "app.version=$(VERSION)" \
	--build-arg APP_NAME=$(APP_NAME) \
	--pull --cache-from=$(IMAGE_NAME) \
	-f Dockerfile .

.PHONY: list-images # list all images with APP_NAME
list-images:
	@docker images --filter=label=app.name=$(APP_NAME)

.PHONY: prune-images # remove all images with APP_NAME
prune-images:
	@docker rmi -f `docker images --filter=label=app.name=$(APP_NAME) -q`

push-image:
	$(call check_defined,VERSION)
	@docker tag $(IMAGE_NAME):$(VERSION) $(IMAGE_NAME):latest
	@docker push $(IMAGE_NAME):$(VERSION)
	@docker push $(IMAGE_NAME):latest

.PHONY: deploy # deploy application with DEPLOY_TO and VERSION
deploy:
	$(call check_defined,VERSION)
	$(call check_defined,DEPLOY_TO)
	@helm --namespace $(NS) \
	upgrade --install $(APP_NAME) $(HELM_REPO_NAME)/$(CHART_NAME) \
	--values ./deployments/configs/$(DEPLOY_TO)/$(APP_NAME).yaml \
	--set image.tag=$(VERSION)

.PHONY: gen # generate all generate commands
gen: gen-wire gen-swagger

.PHONY: gen-wire # generate wire code
gen-wire:
	@wire gen ./...

.PHONY: gen-pb # generate protobuf messages and services
gen-pb:
	@protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./pb/*.proto
	@echo Successfully generated proto

.PHONY: gen-swagger # generate swagger spec
gen-swagger:
	@swag init -g cmd/$(APP_NAME)/main.go --parseInternal -o api/docs

.PHONY: migrate-up # execute migrate up with scripts/migrations
migrate-up:
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migrations up

.PHONY: migrate-down # execute migrate down with script/migrations
migrate-down:
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migrations down