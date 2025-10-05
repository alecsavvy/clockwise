default:

buf-lint:
	@echo "linting protobuf schemas..."
	buf lint

buf-generate:
	@echo "generating protobuf code..."
	buf generate

buf-postprocessors:
	@echo "protobuf postprocessors..."
	protoc-gen-ddex --go-package-prefix github.com/alecsavvy/clockwise/api/ddex ./api/ddex

buf-clean:
	@echo "cleaning generated protobuf code..."
	rm -rf api

gen: buf-clean buf-generate buf-postprocessors
