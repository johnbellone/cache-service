PKG=github.com/johnbellone/cache-service
DATE?= $(shell date +%FT%T%z)
COMMIT?=$(shell git describe --tags --always --dirty)

all: build

proto:
	protoc -I/usr/local/include -I. --micro_out=. --go_out=. protobuf/*.proto

deps:
	go get ./...

build: deps
	go build -tags release \
		-ldflags '-X $(PKG)/main.GitCommit=$(COMMIT) -X $(PKG)/main.BuildDate=$(DATE)' \
		-o bin/cache-service \
		src/*.go
