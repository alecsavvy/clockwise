default:
	make down
	make up

up:
	docker compose up --build -d

down:
	docker compose down

init:
	cometbft init --home ./cmt-home

init-testnet:
	cometbft testnet --n=4 --v=3 --o=./testnet-home --starting-ip-address 192.167.10.2
