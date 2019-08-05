REGISTRY?=registry.gitlab.com/everest-mobile-seiii-se/safebox/safebox-api
APP_VERSION?=latest
BUILD?=go build -ldflags="-w -s"

.PHONY: proto

default: build

build: build-api build-cron

build-api: format lint
	$(BUILD) -o api-server main.go

build-cron: format lint
	$(BUILD) -o api-cron cmd/cron/main.go

proto:
	protoc -I proto services.proto --go_out=plugins=grpc:services

dump:
	./scripts/dump_db.sh

deps:
	dep ensure -vendor-only

test:
	go test $$(go list ./... | grep -v /vendor/)

format:
	go fmt $$(go list ./... | grep -v /vendor/)

vet:
	go vet $$(go list ./... | grep -v /vendor/)

lint:
	golint -set_exit_status -min_confidence 0.3 $$(go list ./... | grep -v /vendor/)

registry: registry-build registry-push

registry-build:
	docker build --pull -t $(REGISTRY):$(APP_VERSION) .
	docker build --pull -f docker/cron/Dockerfile -t $(REGISTRY)/cron:$(APP_VERSION) .

registry-pull:
	docker pull $(REGISTRY):$(APP_VERSION)
	docker pull $(REGISTRY)/cron:$(APP_VERSION)

registry-push:
	docker push $(REGISTRY):$(APP_VERSION)
	docker push $(REGISTRY)/cron:$(APP_VERSION)

registry-clear:
	docker image rm -f $(REGISTRY):$(APP_VERSION)
	docker image rm -f $(REGISTRY)/cron:$(APP_VERSION)

stop:
	docker-compose stop

stop-prod:
	docker stack rm app

dev:
	docker-compose build
	docker-compose up -d
	clear
	@echo ""
	@echo "starting command line:"
	@echo "** when finish exist and run: make stop**"
	@echo ""
	docker-compose exec server sh

prod:
	docker stack deploy --compose-file docker-stack.yml app --with-registry-auth
	clear
	@echo ""
	@echo "commands:"
	@echo "- make stop-prod"
	@echo ""