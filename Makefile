build:
	@go generate
	@go mod tidy
	@go build -o bin/cache-service
