.PHONY: run test install build buildlinux lint

all: install

install:
	go install ./cmd/...

build:
	go build -o ./bin/textractor.exe ./cmd/...

buildlinux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/textractor ./cmd/...
	
run:	
	make build && ./bin/textractor.exe

test:
	go test -timeout 30s ./... -count=1

lint:
	golint ./...
