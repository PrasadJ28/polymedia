.PHONY: build run test test-verbose test-file test-cover clean

build:
	@go build -o bin/fs ./cmd/rest_api

run: build
	@./bin/fs

test:
	@go test ./...

test-verbose:
	@go test ./... -v

test-file:
	@go test -v ./$(FILE)

test-cover:
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

clean:
	@rm -rf bin/
	@rm -f coverage.out coverage.html
