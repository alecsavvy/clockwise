default:
	make down
	make init-testnet
	make up

up:
	docker compose up --build -d

down:
	docker compose down
	rm -rf testnet-home

init:
	cometbft init --home ./cmt-home

init-testnet:
	cometbft testnet --n=4 --v=3 --config ./config_template.toml --o=./testnet-home --starting-ip-address 192.167.10.2
