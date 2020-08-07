package service

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/Mau-MR/rpcbackend/pb"
	"github.com/Mau-MR/rpcbackend/sample"
	"github.com/Mau-MR/rpcbackend/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateClient(t *testing.T) {
	t.Parallel()
	clientStore := NewInMemoryClientStore()
	serverAddress := startTestClientServer(t, clientStore, nil)
	clientClient := newTestClient(t, serverAddress)
	client := sample.NewClient()
	req := &pb.CreateClientReq{
		Name:    client.Name,
		Surname: client.Surname,
		Phone:   client.Phone,
		Id:      client.Id,
	}
	res, err := clientClient.CreateClient(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	//check that client is really store on server
	other, err := clientStore.Find(res.Data.Id)
	require.NoError(t, err)
	require.NotNil(t, other)
	//check that the client is the same
	//TODO: requireSameClient() call here to match

}
func requireSameClient(t *testing.T, client1 *pb.Client, client2 *pb.Client) {
	//this is for checking that are the same
	//they are parsed to json just for internal functions of protobuf
	//dont actually let compare if they are message
	json1, err := serializer.ProtobufToJSON(client1)
	require.NoError(t, err)
	json2, err := serializer.ProtobufToJSON(client2)
	require.NoError(t, err)
	require.Equal(t, json2, json1)
}
func startTestClientServer(t *testing.T, clientStore ClientStore, imageStore ImageStore) string {
	clientServer := NewClientServer(clientStore, imageStore)
	grpcServer := grpc.NewServer()
	pb.RegisterClientServiceServer(grpcServer, clientServer)
	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	go grpcServer.Serve(listener) //non block call

	return listener.Addr().String()
}
func TestClientUploadImage(t *testing.T) {
	t.Parallel()
	testImageFolder := "../tmp"
	clientStore := NewInMemoryClientStore()
	imageStore := NewDiskImageStore(testImageFolder)
	client := sample.NewClient()
	err := clientStore.Save(client)
	require.NoError(t, err)
	serverAddress := startTestClientServer(t, clientStore, imageStore)
	clientClient := newTestClient(t, serverAddress)
	imagePath := fmt.Sprintf("%s/client.jpg", testImageFolder)
	file, err := os.Open(imagePath)
	require.NoError(t, err)
	defer file.Close()
	stream, err := clientClient.UploadImage(context.Background())
	require.NoError(t, err)

	imageType := filepath.Ext(imagePath)
	req := &pb.UploadImageReq{
		Data: &pb.UploadImageReq_Info{
			Info: &pb.ImageInfo{
				ClientId:  client.GetId(),
				ImageType: imageType,
			},
		},
	}
	err = stream.Send(req)
	require.NoError(t, err)
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	size := 0
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		size += n
		req := &pb.UploadImageReq{
			Data: &pb.UploadImageReq_ChunkData{
				ChunkData: buffer[:n],
			},
		}
		err = stream.Send(req)
		require.NoError(t, err)
	}
	res, err := stream.CloseAndRecv()
	require.NoError(t, err)
	require.NotZero(t, res.GetId())
	require.Equal(t, uint32(size), res.GetSize())
	savedImagePath := fmt.Sprintf("%s/%s%s", testImageFolder, res.GetId(), imageType)
	require.FileExists(t, savedImagePath)
	require.NoError(t, os.Remove(savedImagePath))
}
func TestClientSearchLaptop(t *testing.T) {
	t.Parallel()
	filter := &pb.ClientFilter{
		Name:    "Mauricio",
		Surname: "Merida",
		Phone:   "7773491106",
	}
	clientStore := NewInMemoryClientStore()
	expectedIDs := make(map[string]bool)
	for i := 0; i < 3; i++ {
		client := sample.NewClient()
		switch i {
		case 0:
			client.Name = "Nicolas"
		case 1:
			client.Surname = "Peredo"
		case 2:
			client.Phone = "829"
		}
		err := clientStore.Save(client)
		require.NoError(t, err)
	}
	serverAddress := startTestClientServer(t, clientStore, nil)

	clientClient := newTestClient(t, serverAddress)
	req := &pb.SearchClientReq{Filter: filter}
	stream, err := clientClient.SearchClient(context.Background(), req)
	require.NoError(t, err)
	found := 0
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		require.Contains(t, expectedIDs, res.GetClient().Id)
		found += 1
	}
	require.Equal(t, len(expectedIDs), found)

}
func newTestClient(t *testing.T, serverAddress string) pb.ClientServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewClientServiceClient(conn)
}
