GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

#Project variables
PROJECTDIR=$(shell pwd)
GOBIN=$(PROJECTDIR)/bin

BINARYNAME=Akira

all: test build 

test:

build:
	@GOBIN=$(GOBIN) $(GOBUILD) -o $(GOBIN)/$(BINARYNAME) -v 

run: build 
	@$(GOBIN)/$(BINARYNAME) -t $(AKIRA_TOKEN)
