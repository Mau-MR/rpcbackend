package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/Mau-MR/rpcbackend/pb"
	"github.com/Mau-MR/rpcbackend/sample"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddres := flag.String("address", "", "the server addres")
	flag.Parse()
	log.Printf("dial server: %s", *serverAddres)

	conn, err := grpc.Dial(*serverAddres, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	clientClient := pb.NewClientServiceClient(conn)

	client := sample.NewClient()
	req := &pb.CreateClientReq{
		Name:    client.Name,
		Surname: client.Surname,
		Phone:   client.Phone,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := clientClient.CreateClient(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			//not to important
			log.Print("client already exist")
		} else {
			log.Fatal("cannot create client: ", err)
		}
		return
	}
	log.Printf("Created client with phone: %s", res.Data.Phone)

}
