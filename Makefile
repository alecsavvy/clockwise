default:
	make down
	make up

up:
	docker compose up --build -d

down:
	docker compose down

init:
	cometbft init --home ./cmt-home
