default:
	make down
	make gen
	make up

up:
	cd infra && docker compose up --build -d

down:
	cd infra && docker compose down
	rm -rf testnet-home

init:
	cometbft init --home ./cmt-home

init-testnet:
	cometbft testnet --n=3 --v=4 --config ./infra/config_template.toml --o=./testnet-home --starting-ip-address 192.167.10.2

gen:
	cd core/db && sqlc generate
	go generate ./...
	protoc --go_out=./protocol --go-grpc_out=./protocol ./protocol/protocol.proto
	make init-testnet
	go mod tidy

deps:
	brew install protobuf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/cometbft/cometbft/cmd/cometbft@v0.38.9

mosh:
	docker start moshpit
	open http://localhost:8080/

chill:
	docker stop moshpit
