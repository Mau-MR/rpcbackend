package service

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"

	"github.com/Mau-MR/rpcbackend/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//meximun 1 megabyte
const maxImageSize = 1 << 20

//ClientService is the server that provices laptop services
type ClientServer struct {
	clientStore ClientStore
	imageStore  ImageStore
}

func NewClientServer(clientStore ClientStore, imageStore ImageStore) *ClientServer {
	return &ClientServer{clientStore, imageStore}
}

func (server *ClientServer) UploadImage(stream pb.ClientService_UploadImageServer) error {
	req, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot receive image info"))
	}
	clientId := req.GetInfo().GetClientId()
	imageType := req.GetInfo().GetImageType()
	log.Printf("receive an upload image request for client %s, with image type %s", clientId, imageType)
	client, err := server.clientStore.Find(clientId)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot find client:%v ", err))
	}
	if client == nil {
		return logError(status.Errorf(codes.NotFound, "client %s doesn't exit", clientId))
	}
	imageData := bytes.Buffer{}
	imageSize := 0
	for {
		if err := contextErr(stream.Context()); err != nil {
			return err
		}
		log.Printf("waiting to receive more data...")
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}
		chunk := req.GetChunkData()
		size := len(chunk)
		log.Printf("received chunk with size: %d", size)

		imageSize += size
		if imageSize > maxImageSize {
			return logError(status.Errorf(codes.InvalidArgument, "image is to large %d > %d", imageSize, maxImageSize))
		}
		_, err = imageData.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}
	imageId, err := server.imageStore.Save(clientId, imageType, imageData)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot write data to the store: %v", err))
	}
	res := &pb.UploadImageRes{
		Id:   imageId,
		Size: uint32(imageSize),
	}
	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response %v", err))
	}
	log.Printf("saved image with id: %s, size: %d", imageId, imageSize)

	return nil
}
func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
func contextErr(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceed"))
	default:
		return nil
	}
}

//CreatCreateClient is a unary rpc to create a newClient
func (server *ClientServer) CreateClient(ctx context.Context, req *pb.CreateClientReq) (*pb.ClientRes, error) {
	client := &pb.Client{
		Name:    req.Name,
		Surname: req.Surname,
		Phone:   req.Phone,
		Id:      req.Id,
	}
	log.Printf("receive a createclient request with id: %s", client.Id)
	if len(client.Id) > 0 {
		_, err := uuid.Parse(client.Id)
		if err != nil {
			return nil, logError(status.Errorf(codes.InvalidArgument, "client ID is not a valid UUID: %v", err))
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, logError(status.Errorf(codes.Internal, "cannot generate a new client ID: %v", err))
		}
		client.Id = id.String()
	}
	//save  client to store
	if err := contextErr(ctx); err != nil {
		return nil, err
	}
	err := server.clientStore.Save(client)
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
func (server *ClientServer) SearchClient(req *pb.SearchClientReq, stream pb.ClientService_SearchClientServer) error {
	filter := req.GetFilter()
	log.Printf("receive a search client request with filter: %v", filter)
	err := server.clientStore.Search(stream.Context(), filter, func(client *pb.Client) error {
		res := &pb.SearchClientRes{Client: client}
		err := stream.Send(res)
		if err != nil {
			return err
		}
		log.Printf("send client with phone: %v", client.GetPhone())
		return nil

	},
	)
	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}
	return nil
}
