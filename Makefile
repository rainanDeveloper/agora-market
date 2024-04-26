.PHONY: default run build test

default: run-with-docs

build:
	@go build -o bin/server cmd/server/main.go

run:
	@go run cmd/server/main.go

run-with-docs:
	@swag init
	@go run cmd/server/main.go