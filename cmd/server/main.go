package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Mau-MR/rpcbackend/pb"
	"github.com/Mau-MR/rpcbackend/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	tokenDuration = 5 * time.Minute
	secretKey     = "secretjaja"
)

func seedUsers(userStore service.UserStore) error {
	err := createUser(userStore, "admin", "secret", "admin")
	if err != nil {
		return err
	}
	return createUser(userStore, "user1", "secret", "user")

}
func accesibleRoles() map[string][]string {
	const clientServicePath = "/pb.ClientService/"
	return map[string][]string{
		clientServicePath + "CreateClient": {"admin"},
		clientServicePath + "UploadImage":  {"user", "admin"},
	}

}
func createUser(userStore service.UserStore, username, password, role string) error {
	user, err := service.NewUser(username, password, role)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}
func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	userStore := service.NewInMemoryUserStore()
	err := seedUsers(userStore)
	if err != nil {
		log.Fatal("cannot generete seed users")
	}
	jwtManager := service.NewJWTManager(secretKey, tokenDuration)
	authServer := service.NewAuthServer(userStore, jwtManager)

	clientStore := service.NewInMemoryClientStore()
	imageStore := service.NewDiskImageStore("img")
	clientServer := service.NewClientServer(clientStore, imageStore)

	interceptor := service.NewAuthInterceptor(jwtManager, accesibleRoles())
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	pb.RegisterAuthServiceServer(grpcServer, authServer)
	pb.RegisterClientServiceServer(grpcServer, clientServer)
	reflection.Register(grpcServer)

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
