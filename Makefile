build:
	@go mod tidy
	@go build -o bin/cache-service

proto:
	@go generate
