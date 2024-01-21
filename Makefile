include .dev.env
export

PROJECTNAME=$(shell basename "$(PWD)")

# Go related variables.
GOBASE=$(shell pwd)
#GOPATH="$(GOBASE)/vendor:$(GOBASE)"
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard $(GOBASE)/cmd/$(PROJECTNAME)/*.go)

## clean: Clean build files. Runs `go clean` internally.
clean:
	@echo "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean
	rm -f $(GOBIN)/$(PROJECTNAME)

tidy:
	@echo "  >  Update dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go mod tidy

build:
	@echo "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

run:
	@echo "  >  Start $(PROJECTNAME)..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go run $(GOFILES)

