BINARIES=$$(go list ./cmd/...)
TESTABLE=$$(go list ./...)

all: test build

deps:
	@go get -u && go mod tidy
.PHONY: deps

build:
	@go install -v $(BINARIES)
.PHONY: build

test:
	@go test -v $(TESTABLE)
.PHONY: test

local: test build
	@./scripts/local.sh
.PHONY: local
