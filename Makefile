gen:
	protoc --proto_path=proto proto/*.proto -I. --go_out=plugins=grpc:/home/mawi/gocode/src
clean:
	rm pb/*.go
server:
	go run cmd/server/main.go -port 8080
client:
	go run cmd/client/main.go -address 0.0.0.0:8080
test:
	go test -cover -race ./...

.PHONY: gen clean server test client
