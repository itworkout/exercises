.PHONY: tools
tools: 
	$(info # Generate tools in ./bin ...)
	@cd tools && go mod download && go generate -tags tools

.PHONY: lint
lint:
	$(info # Running lint ...)
	@./bin/golangci-lint run --config=.golangci.yml ./...

.PHONY: format
format:
	$(info # Formatting ...)
	@./bin/goimports -w ./src

.PHONY: test
test:
	$(info # Testing ...)
	@go test -cover -race ./...
