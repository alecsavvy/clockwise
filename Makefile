default:

buf-lint:
	@echo "linting protobuf schemas..."
	buf lint

buf-generate:
	@echo "generating protobuf code..."
	buf generate

buf-clean:
	@echo "cleaning generated protobuf code..."
	rm -rf api/

buf-regen: buf-clean buf-generate
	@echo "regenerated protobuf code"

proto-clean: buf-clean

proto-gen: buf-generate

proto: buf-regen
