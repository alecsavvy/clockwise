GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
BINARY_NAME=myapp

all: build run

build:
	make gen
	$(GOBUILD) -o $(BINARY_NAME) main.go

run:
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

dev:
	$(GORUN) main.go

deps:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


gen:
	cd ./db/sql && sqlc generate
	cd ./api/proto && protoc --proto_path=. --go_out=../../ --go-grpc_out=../../ api.proto
