package main

import (
	"flag"
	"github.com/Mau-MR/rpcbackend/client"
	"github.com/Mau-MR/rpcbackend/pb"
	"github.com/Mau-MR/rpcbackend/sample"
	"time"

	"google.golang.org/grpc"
	"log"
)

func testCreateClient(clientClient client.ClientClient) {
	clientClient.CreateClient(sample.NewClient())
}
func testSearchClient(clientClient client.ClientClient) {
	for i := 0; i < 10; i++ {
		clientClient.CreateClient(sample.NewClient())
	}
	filter := &pb.ClientFilter{
		Name:    "Mauricio",
		Surname: "Merida",
		Phone:   "7773491106",
	}
	clientClient.SearchClient(filter)

}
func testUploadImage(clientClient client.ClientClient) {
	client := sample.NewClient()
	clientClient.CreateClient(client)
	clientClient.UploadImage(client.GetId(), "tmp/client.jpg")
}
func testCreateAccount(authClient client.AuthClient) {
	authClient.CreateAccount()

}

const (
	username        = "Mawi"
	password        = "secret"
	refreshDuration = 30 * time.Second
)

func authMethods() map[string]bool {

	const clientServicePath = "/pb.ClientService/"
	return map[string]bool{
		clientServicePath + "CreateClient": true,
		clientServicePath + "UploadImage":  true,
	}

}
func main() {
	serverAddres := flag.String("address", "", "the server addres")
	flag.Parse()
	log.Printf("dial server: %s", *serverAddres)

	cc1, err := grpc.Dial(*serverAddres, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	authClient := client.NewAuthClient(cc1, username, password)
	interceptor, err := client.NewAuthInterceptor(authClient, authMethods(), refreshDuration)
	if err != nil {
		log.Fatal("cannot create auth interceptor: ", err)
	}
	cc2, err := grpc.Dial(
		*serverAddres,
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(interceptor.Unary()),
		grpc.WithStreamInterceptor(interceptor.Stream()),
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	clientClient := client.NewClientClient(cc2)
	testCreateClient(*clientClient)
	testCreateAccount(*authClient)
	//testUploadImage(clientClient)

}
