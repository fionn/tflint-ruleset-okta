.PHONY: default
default: build

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	go build

.PHONY: install
install: build
	mkdir -p ~/.tflint.d/plugins
	cp ./tflint-ruleset-okta ~/.tflint.d/plugins
