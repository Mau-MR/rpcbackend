package client

import (
	"context"
	"time"

	"github.com/Mau-MR/rpcbackend/pb"
	"google.golang.org/grpc"
)

type AuthClient struct {
	service  pb.AuthServiceClient
	username string
	password string
}

//NewAuthClient returns a new auth client
func NewAuthClient(cc *grpc.ClientConn, username, password string) *AuthClient {
	service := pb.NewAuthServiceClient(cc)
	return &AuthClient{service, username, password}
}

//Login logs in user and returns access token
func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}
	res, err := client.service.Login(ctx, req)
	if err != nil {
		return "", err
	}
	return res.GetAccessToken(), nil

}
func (client *AuthClient) CreateAccount() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.CreateAccountReq{
		User:     "Mawi",
		Name:     "Mauricio Eulalio Merida Rivera",
		Phone:    "7773491106",
		Bussines: "Thea",
		Password: "Secreto",
	}
	res, err := client.service.CreateAccount(ctx, req)
	if err != nil {
		return "", err
	}
	return res.GetBid(), nil

}
