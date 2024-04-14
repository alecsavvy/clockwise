run:
	air

deps:
	go install github.com/cosmtrek/air@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

gen:
	cd ./db/sql && sqlc generate