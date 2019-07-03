GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -v

#Project variables
PROJECTDIR=$(shell pwd)
TESTDIR=$(PROJECTDIR)/test
GOBIN=$(PROJECTDIR)/bin

BINARYNAME=Akira

all: build test

.PHONY: test

test: 
	@GOBIN=$(GOBIN) $(GOTEST) $(TESTDIR)

build:
	@GOBIN=$(GOBIN) $(GOBUILD) -o $(GOBIN)/$(BINARYNAME) -v 

run: build 
	@$(GOBIN)/$(BINARYNAME) -t $(AKIRA_TOKEN)
