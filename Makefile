PROTO_SRC=./proto
GEN_SRC=./gen
CLIENT_SRC=./client
SERVER_SRC=./server
PROXY_SRC=./proxy


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

# example: three servers, two cients
PROXY_PORT=:50051
SERVER1_PORT=:50052
SERVER2_PORT=:50053
SERVER3_PORT=:50054

run_example: | protoc
	go run $(SERVER_SRC) -- $(SERVER1_PORT) server1 
	go run $(SERVER_SRC) -- $(SERVER2_PORT) server2
	go run $(SERVER_SRC) -- $(SERVER3_PORT) server3
	go run $(PROXY_SRC) -- $(PROXY_PORT) "localhost$(SERVER1_PORT)" "localhost$(SERVER2_PORT)" "localhost$(SERVER3_PORT)"
	go run $(CLIENT_SRC) -- client1 "localhost$(PROXY_PORT)"
	go run $(CLIENT_SRC) -- client2 "localhost$(PROXY_PORT)"