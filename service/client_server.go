package service

import (
	"context"
	"errors"
	"log"

	"github.com/Mau-MR/rpcbackend/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//ClientService is the server that provices laptop services
type ClientServer struct {
	Store ClientStore
}

func NewClientServer(store ClientStore) *ClientServer {
	return &ClientServer{store}
}

//CreatCreateClient is a unary rpc to create a newClient
func (server *ClientServer) CreateClient(ctx context.Context, req *pb.CreateClientReq) (*pb.ClientRes, error) {

	client := &pb.Client{
		Name:    req.Name,
		Surname: req.Surname,
		Phone:   req.Phone,
	}
	log.Printf("receive a createclient request with phone: %s", client.Phone)
	//save  client to store
	if ctx.Err() == context.Canceled {
		log.Print("request is canceled")
		return nil, status.Error(codes.Canceled, "request is canceled")
	}
	if ctx.Err() == context.DeadlineExceeded {
		log.Print("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceed")
	}
	err := server.Store.Save(client)
	if err != nil {
		//also can use for other time codes.InvalidArgument
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			log.Printf("client already exist")
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot save client to store")
	}
	res := &pb.ClientRes{
		Success: true,
		Data:    client,
	}
	log.Printf("client succes write to store")
	//TODO: Make the db
	return res, nil

}
func (server *ClientServer) UpdateClient(context.Context, *pb.UpdateClientReq) (*pb.ClientRes, error) {
	return &pb.ClientRes{Success: true}, nil

}
