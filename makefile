PROJECTNAME=$(shell basename $$(pwd))
SRC := $(PROJECTNAME)/src
MODULE := findlinks
GOBIN = $(PROJECTNAME)/bin
EXECUTABLE := findlinks1.exe
all:
go-build:
    @echo "  >  Building binary..."
    @GOBIN=$(GOBIN) go build -o $(GOBIN)/$(PROJECTNAME) $(EXECUTABLE)