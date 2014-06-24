export GOPATH=$(shell pwd)

default: fmt build test

fmt:
	go fmt

build:
	go build

test:
	go test
