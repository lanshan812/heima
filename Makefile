SHELL := /bin/bash
BASEDIR = $(shell pwd)
BINARY = patent

build-test:export Patent_ENV="test"
build-test:
	go clean
	@echo "current Patent_ENV env is" $(Patent_ENV)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}
build-prod:export Patent_ENV="production"
build-prod:
	go clean
	@echo "current Patent_ENV env is" $(Patent_ENV)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

clean:
	rm -rf patent
	find . -name "[._]*.s[a-w][a-z]" | xargs rm -f {}
gotool:
	gofmt -w .
	go vet .

help:
	@echo "make build - compile the source code and build binary"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"

.PHONY: clean gotool ca help