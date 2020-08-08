package client

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Mau-MR/rpcbackend/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClientClient struct {
	service pb.ClientServiceClient
}

func NewClientClient(cc *grpc.ClientConn) *ClientClient {
	service := pb.NewClientServiceClient(cc)
	return &ClientClient{service}
}

func (clientClient *ClientClient) CreateClient(client *pb.Client) {

	req := &pb.CreateClientReq{
		Name:    client.Name,
		Surname: client.Surname,
		Phone:   client.Phone,
		Id:      client.Id,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := clientClient.service.CreateClient(ctx, req)
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
func (clientClient *ClientClient) SearchClient(filter *pb.ClientFilter) {
	log.Print("search filter: ", filter)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req := &pb.SearchClientReq{Filter: filter}
	stream, err := clientClient.service.SearchClient(ctx, req)
	if err != nil {
		log.Fatal("cannot search client: ", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("cannot receive response:", err)
		}
		client := res.GetClient()
		log.Print("found, client")
		log.Print("name: ", client.Name)
		log.Print("surname: ", client.Surname)
		log.Print("phone: ", client.Phone)
	}

}
func (clientClient *ClientClient) UploadImage(clientId string, imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stream, err := clientClient.service.UploadImage(ctx)
	if err != nil {
		log.Fatal("cannot uploadImage")
	}
	req := &pb.UploadImageReq{
		Data: &pb.UploadImageReq_Info{
			Info: &pb.ImageInfo{
				ClientId:  clientId,
				ImageType: filepath.Ext(imagePath),
			},
		},
	}
	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send image info to the server: ", err, stream.RecvMsg(nil))
	}
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer", err)
		}
		req := &pb.UploadImageReq{
			Data: &pb.UploadImageReq_ChunkData{
				ChunkData: buffer[:n],
			},
		}
		err = stream.Send(req)
		if err != nil {
			log.Fatal("cannot send chunk to the server: ", err, stream.RecvMsg(nil))
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response")
	}
	log.Printf("immage uploaded with id: %s, size: %d", res.GetId(), res.GetSize())

}
