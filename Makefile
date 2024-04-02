BINARY_NAME=api

build:
	@GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME} ./cmd/api

run: build
	@./bin/${BINARY_NAME}

clean:
	@go clean
	@rm bin/${BINARY_NAME}

format:
	@gofmt -s -w .
	@goimports .
	@echo "\n fumpt and imports -> done!"

test:
	@go test -v ./...

test_coverage:
	@go test ./... -coverprofile=coverage.out

lint:
	@golangci-lint run

