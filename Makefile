PROTO_SRC=./proto
GEN_SRC=./gen
CLIENT_SRC=./client
SERVER_SRC=./server


protoc:
	rm -rf $(GEN_SRC) && mkdir $(GEN_SRC)
	protoc --go_out=$(GEN_SRC) --go_opt=paths=source_relative \
    --go-grpc_out=$(GEN_SRC) --go-grpc_opt=paths=source_relative \
    $(PROTO_SRC)/example.proto
	mv $(GEN_SRC)/proto/* $(GEN_SRC)
	rm -r $(GEN_SRC)/proto

build: | protoc
	go mod tidy
	go build $(CLIENT_SRC)
	go build $(SERVER_SRC)

run_server: | protoc
	go run $(SERVER_SRC)

run_client: | protoc
	go run $(CLIENT_SRC)

build_docker:
	docker build -t grpc_example -f server.Dockerfile .

run_example: build_docker
	docker-compose up