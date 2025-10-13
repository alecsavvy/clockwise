run:
	@go run ./cmd/clockwised/main.go

buf-lint:
	@echo "linting protobuf schemas..."
	buf lint

buf-generate:
	@echo "generating protobuf code..."
	buf generate

buf-clean:
	@echo "cleaning generated protobuf code..."
	rm -rf api

gen: buf-clean buf-generate
