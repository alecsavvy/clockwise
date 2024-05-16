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
	cometbft testnet --n=4 --v=3 --config ./infra/config_template.toml --o=./testnet-home --starting-ip-address 192.167.10.2

gen:
	make init-testnet
	go generate ./...
	cd core/db && sqlc generate
	go mod tidy
