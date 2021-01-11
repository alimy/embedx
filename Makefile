GOFMT ?= gofmt -s -w
GOFILES := $(shell find . -name "*.go" -type f)

.PHONY: test
test:
	@go test ./...

.PHONY: fmt
fmt:
	$(GOFMT) $(GOFILES)
