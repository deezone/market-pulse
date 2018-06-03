# Set up
APPLICATION=forex-clock
BIN=forex-clock
DOCSCMD=aglio
GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get

# Build and test during normal execution
all: build test

build: get
	@echo "Building binary..."
	$(GOBUILD) -o dist/$(BIN)

docs: docs-deps
	@echo "Generating API documentation"
	npm run-script generate-api-docs
	@echo "Documentation generated, open docs/api-output.html via a browser."

docs-deps:
	@echo "Installing documentation dependencies..."
	@echo "Installing $(DOCSCMD)..." && npm install

get:
	@echo "Getting external packages..."
	$(GOGET) -u -t ./...

release: build
	@echo "Releasing version $V..."
	pops release -a $(APPLICATION) -v $V -d .

run: build
	@echo "Running built binary..."
	dist/$(BIN)

run-prod:
	@echo "Running built binary..."
	dist/$(BIN)

.PHONY: all build docs docs-deps get release run run-prod
