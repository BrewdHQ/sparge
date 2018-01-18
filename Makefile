BINARY=sparge
VERSION=1.0.0

glide-update:
	glide cc; glide update

build:
	go build -o $(BINARY) main.go

build-linux:
	GOOS=linux CGO_ENABLED=0 go build -o $(BINARY) main.go

docker-build: glide-update build-linux
	docker build -t brewd/sparge:$(VERSION) .
	docker image tag brewd/sparge:$(VERSION) brewd/sparge:latest

.PHONY: glide-update build build-linux docker-build