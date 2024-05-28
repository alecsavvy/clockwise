default:
	make down
	make gen
	make up
	open http://localhost:8080/

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
	make init-testnet
	go generate ./...
	go mod tidy
