GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test -v

#Project variables
PROJECTDIR=$(shell pwd)
TESTDIR=$(PROJECTDIR)/test
GOBIN=$(PROJECTDIR)/bin

DOCKERCMD=docker

BINARYNAME=Akira
IMAGENAME=akira

all: build test

.PHONY: test

test: 
	@GOBIN=$(GOBIN) $(GOTEST) $(TESTDIR)

build:
	@GOBIN=$(GOBIN) $(GOBUILD) -o $(GOBIN)/$(BINARYNAME) -v 

run: build 
	@$(GOBIN)/$(BINARYNAME) -t $(AKIRA_TOKEN)

docker: 
	@$(DOCKERCMD) build -t $(IMAGENAME)  .

run-docker: docker
	@$(DOCKERCMD) run -e "AKIRA_DISCORD_TOKEN=$(AKIRA_TOKEN)" $(IMAGENAME)