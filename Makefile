BINARY=sparge
VERSION=1.1.0
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

glide-update:
	glide cc; glide update

build:
	go build -o $(BINARY) main.go

build-linux:
	GOOS=linux CGO_ENABLED=0 go build $(LDFLAGS) -o $(BINARY) main.go

docker-build: clean glide-update build-linux
	docker build -t brewd/sparge:$(VERSION) .
	docker image tag brewd/sparge:$(VERSION) brewd/sparge:latest

clean:
	rm $(BINARY)

.PHONY: glide-update build build-linux docker-build clean