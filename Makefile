LOCAL_BIN=$(CURDIR)/bin
PROJECT_NAME=etherium-proxy

.PHONY: build
build: test
	go build -v -o $(LOCAL_BIN)/$(PROJECT_NAME) ./cmd

.PHONY: run
run:
	go run cmd/main.go

.PHONY: deps
deps:
	go mod tidy
	go get google.golang.org/protobuf/cmd/protoc-gen-go \
    	google.golang.org/grpc/cmd/protoc-gen-go-grpc
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: generate
generate:
	protoc -I ./pb/ \
	  -I . \
	  --grpc-gateway_out ./pkg \
	  --go_out ./pkg \
	  --go-grpc_out ./pkg \
      --grpc-gateway_opt logtostderr=true \
      --grpc-gateway_opt paths=source_relative \
      api/etherium-proxy.proto

.PHONY: generateMocks
generateMocks:
	mockery --dir ./internal/pkg/ethservice --output ./internal/pkg/ethservice/mocks --name=EthClient

.PHONY: lint
lint:
	$(LOCAL_BIN)/golangci-lint run ./...

.PHONY: test
test:
	go test ./...