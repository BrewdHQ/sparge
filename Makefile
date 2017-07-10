BINARY=sparge
VERSION=0.3.0

build:
	go build -o $(BINARY) main.go

build-linux:
	GOOS=linux CGO_ENABLED=0 go build -o $(BINARY) main.go

docker-build: build-linux
	docker build -t brewd/sparge:$(VERSION) .
	docker image tag brewd/sparge:$(VERSION) brewd/sparge:latest

.PHONY: build build-linux docker-build