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

proto: proto-go proto-dart

proto-deps:
	go get github.com/gogo/protobuf/protoc-gen-gofast
	flutter pub global activate protoc_plugin

proto-dart:
	rm -rf output/dart && mkdir -p output/dart
	protoc -Iproto --dart_out=grpc:./output/dart services.proto

proto-go:
	protoc -Iproto --go_out=plugins=grpc:./services services.proto

gomobile:
	GO111MODULE=off go get golang.org/x/mobile/cmd/gomobile
	gomobile init

mobile: mobile-android mobile-ios

mobile-android:
	rm -rf output/android && mkdir -p output/android
	gomobile bind -ldflags="-w -s" -target=android -o ./output/android/Safebox.aar safebox.jerson.dev/api/mobile

mobile-ios:
	rm -rf output/ios && mkdir -p output/ios
	gomobile bind -ldflags="-w -s" -target=ios -o ./output/ios/Safebox.framework safebox.jerson.dev/api/mobile

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