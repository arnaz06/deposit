TEST_OPTS := -covermode=atomic $(TEST_OPTS)


# Linter
.PHONY: lint-prepare
lint-prepare:
	@echo "Installing golangci-lint"
	@wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.31.0

.PHONY: lint
lint:
	golangci-lint run ./...

# Run http server
.PHONY: run-http
run-http:
	@cd cmd/http && go run . http

# Run proccesor server
.PHONY: run-proccesor
run-proccesor:
	@cd cmd/proccesor && go run . proccesor


.PHONY: vendor
vendor: go.mod go.sum
	@GO111MODULE=on go get ./...

.PHONY: test
test: vendor
	GO111MODULE=on go test $(TEST_OPTS) ./...
