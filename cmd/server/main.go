package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Mau-MR/rpcbackend/pb"
	"github.com/Mau-MR/rpcbackend/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	clientServer := service.NewClientServer(service.NewInMemoryClientStore())
	grpcServer := grpc.NewServer()
	pb.RegisterClientServiceServer(grpcServer, clientServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start sever: ", err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
