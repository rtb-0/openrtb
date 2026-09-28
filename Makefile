.PHONY: fmt tidy lint test bench help

.DEFAULT_GOAL := help

export GOWORK := off

fmt: ## gofmt -w
	gofmt -w $$(go list -f '{{.Dir}}' ./...)

tidy: ## go mod tidy
	go mod tidy

lint: ## go vet ./...
	go vet ./...

test: ## go test ./...
	go test ./...

bench: ## benchmarks only, with memory stats
	go test -bench=. -benchmem -run=^$$ ./...

help: ## show this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-8s %s\n", $$1, $$2}' $(MAKEFILE_LIST)
