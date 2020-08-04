gen:
	protoc --proto_path=proto proto/*.proto -I. --go_out=plugins=grpc:/home/mawi/gocode/src
clean:
	rm pb/*.go
run:
	go run main.go
test:
	go test -cover -race ./...
