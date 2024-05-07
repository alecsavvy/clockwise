default:
	make down
	make init-testnet
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
	go generate ./...