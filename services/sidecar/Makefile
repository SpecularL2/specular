GO_CMD=go
GO_BUILD=$(GO_CMD) build
CMD_PATH=./cmd/sidecar/main.go

.PHONY: help # Help - list of targets with descriptions
help:
	@echo ''
	@echo 'Usage:'
	@echo ' ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@grep '^.PHONY: .* #' Makefile | sed 's/\.PHONY: \(.*\) # \(.*\)/ \1\t\2/' | expand -t20

.PHONY: wire-generate # Generate Wire bindings
wire-generate:
	cd internal/service/di ;\
	wire

.PHONY: lint-test # Run lint-tests
lint-test:
	go mod tidy
	goimports -local github.com/specularL2/specular/services/sidecar -w .
	go fmt ./...
	@golangci-lint -v run ./...
	@test -z "$$(golangci-lint run ./...)"

.PHONY: lint # Run lint-tests (alias)
lint: lint-test
