SRC := $(shell git ls-files *.go)
TARGET := tflint-ruleset-okta

.PHONY: build
build: $(TARGET)

$(TARGET): $(SRC) go.mod go.sum
	@go build

.PHONY: test
test:
	@go test -v ./...

.PHONY: test-coverage
test-coverage: coverage.out
	@go tool cover -func=$<

coverage.out: $(SRC)
	@go test -covermode=count -coverprofile=$@ ./...

.PHONY: clean
clean:
	@rm -f $(TARGET) coverage.out
