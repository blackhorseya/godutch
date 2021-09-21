APP_NAME = cashdivider
VERSION = latest
PROJECT_ID = sean-side
NS = side
DEPLOY_TO = uat

DB_URI = "mysql://cashdivider:changeme@tcp(localhost:3306)/cashdivider?charset=utf8mb4&parseTime=True&loc=Local"

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
	@curl -XPOST 'https://goreportcard.com/checks' --data 'repo=github.com/blackhorseya/cashdivider'

.PHONY: test-unit # execute unit test
test-unit:
	@sh $(shell pwd)/scripts/go.test.sh

.PHONY: build-image # build docker image with APP_NAME and VERSION
build-image:
	@docker build -t $(APP_NAME):$(VERSION) \
	--label "app.name=$(APP_NAME)" \
	--label "app.version=$(VERSION)" \
	--build-arg APP_NAME=$(APP_NAME) \
	--pull \
	-f Dockerfile .

.PHONY: list-images # list all images with APP_NAME
list-images:
	@docker images --filter=label=app.name=$(APP_NAME)

.PHONY: prune-images # remove all images with APP_NAME
prune-images:
	@docker rmi -f `docker images --filter=label=app.name=$(APP_NAME) -q`

.PHONY: tag-image # tag image with gcr.io
tag-image:
	@docker tag $(APP_NAME):$(VERSION) gcr.io/$(PROJECT_ID)/$(APP_NAME):$(VERSION)

.PHONY: push-image # push image to gcr.io
push-image:
	@docker push gcr.io/$(PROJECT_ID)/$(APP_NAME):$(VERSION)

.PHONY: install-db # install database with APP_NAME and DEPLOY_TO
install-db:
	@helm --namespace $(NS) upgrade --install $(APP_NAME)-db bitnami/mariadb \
	--values ./deployments/configs/$(DEPLOY_TO)/db.yaml

.PHONY: deploy # deploy application with DEPLOY_TO and VERSION
deploy:
	@helm --namespace $(NS) \
	upgrade --install $(APP_NAME) ./deployments/$(APP_NAME) \
	--values ./deployments/configs/$(DEPLOY_TO)/$(APP_NAME).yaml \
	--set image.tag=$(VERSION)

.PHONY: gen # generate all generate commands
gen: gen-wire gen-pb gen-swagger

.PHONY: gen-wire # generate wire code
gen-wire:
	@wire gen ./...

.PHONY: gen-pb # generate protobuf messages and services
gen-pb:
	@protoc --go_out=plugins=grpc,paths=source_relative:. ./internal/pkg/entity/**/*.proto

.PHONY: gen-swagger # generate swagger spec
gen-swagger:
	@swag init -g cmd/$(APP_NAME)/main.go --parseInternal -o api/docs

.PHONY: migrate-up # execute migrate up with scripts/migrations
migrate-up:
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migrations up

.PHONY: migrate-down # execute migrate down with script/migrations
migrate-down:
	@migrate -database $(DB_URI) -path $(shell pwd)/scripts/migrations down