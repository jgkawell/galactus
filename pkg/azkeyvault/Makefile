PROJECT=azkeyvault
GO := GO15VENDOREXPERIMENT=1 go
PACKAGE_DIRS := $(shell $(GO) list ./... | grep -v /vendor/)
CGO_ENABLED = 0

all: build

.PHONY: build
build:
	CGO_ENABLED=$(CGO_ENABLED) $(GO) build .

test:
	CGO_ENABLED=$(CGO_ENABLED) $(GO) test $(PACKAGE_DIRS) -test.v -cover

fmt:
	@echo $(PACKAGE_DIRS)
	@FORMATTED=`$(GO) fmt $(PACKAGE_DIRS)`

.PHONY: local clean
local:
	@if [ ! -d "bin" ]; then mkdir bin; fi
	CGO_ENABLED=0 GO15VENDOREXPERIMENT=1 go build -ldflags '' -o bin/azkeyvault azkeyvault.go

clean:
	@rm -rf bin

GOLINT := $(GOPATH)/bin/golint
$(GOLINT):
	go get github.com/golang/lint/golint

.PHONY: lint
lint: 
	@( pwd; ${GOLINT}; )
	@( cd service; ${GOLINT}; )
	@( cd token; ${GOLINT}; )
	@( cd model; ${GOLINT}; )
