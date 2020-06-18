build:
	@go mod tidy
	@go build -o bin/cache-service

proto:
	@go generate

release: build
	@zip --junk-paths bin/cache-service README.md LICENSE
