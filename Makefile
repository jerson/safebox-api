REGISTRY?=registry.gitlab.com/pardacho/safebox-api
APP_VERSION?=latest
BUILD?=go build -ldflags="-w -s"

.PHONY: proto

default: build

build: build-api build-cron build-commands build-queue

build-api: format lint
	$(BUILD) -o api-server main.go

build-cron: format lint
	$(BUILD) -o api-cron cmd/cron/main.go

build-commands: format lint
	$(BUILD) -o api-commands cmd/commands/main.go

build-queue: format lint
	$(BUILD) -o api-queue cmd/queue/main.go

proto:
	protoc -I proto services.proto --go_out=plugins=grpc:services

gomobile:
	GO111MODULE=off go get golang.org/x/mobile/cmd/gomobile
	gomobile init

mobile: mobile-android mobile-ios

mobile-android:
	gomobile bind -ldflags="-w -s" -target=android -o Safebox.aar safebox.jerson.dev/api/mobile

mobile-ios:
	gomobile bind -ldflags="-w -s" -target=ios -o Safebox.framework safebox.jerson.dev/api/mobile

dump:
	./scripts/dump_db.sh

deps:
	go mod download

test:
	go test ./...

format:
	go fmt ./...

vet:
	go vet ./...

lint:
	golint -set_exit_status -min_confidence 0.3 ./...

registry: registry-build registry-push

registry-build:
	docker build --pull -t $(REGISTRY):$(APP_VERSION) .
	docker build --pull -f docker/cron/Dockerfile -t $(REGISTRY)/cron:$(APP_VERSION) .
	docker build --pull -f docker/commands/Dockerfile -t $(REGISTRY)/commands:$(APP_VERSION) .
	docker build --pull -f docker/queue/Dockerfile -t $(REGISTRY)/queue:$(APP_VERSION) .

registry-push:
	docker push $(REGISTRY):$(APP_VERSION)
	docker push $(REGISTRY)/cron:$(APP_VERSION)
	docker push $(REGISTRY)/commands:$(APP_VERSION)
	docker push $(REGISTRY)/queue:$(APP_VERSION)

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