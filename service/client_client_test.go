package service

import (
	"context"
	"net"
	"testing"

	"github.com/Mau-MR/rpcbackend/pb"
	"github.com/Mau-MR/rpcbackend/sample"
	"github.com/Mau-MR/rpcbackend/serializer"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()
	clientServer, serverAddress := startTestClientServer(t)
	clientClient := newTestClient(t, serverAddress)
	client := sample.NewClient()
	req := &pb.CreateClientReq{
		Name:    client.Name,
		Surname: client.Surname,
		Phone:   client.Phone,
	}
	res, err := clientClient.CreateClient(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	//check that client is really store on server
	other, err := clientServer.Store.Find(res.Data.Phone)
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
func startTestClientServer(t *testing.T) (*ClientServer, string) {
	clientServer := NewClientServer(NewInMemoryClientStore())
	grpcServer := grpc.NewServer()
	pb.RegisterClientServiceServer(grpcServer, clientServer)
	listener, err := net.Listen("tcp", ":4567")
	require.NoError(t, err)
	go grpcServer.Serve(listener) //non block call

	return clientServer, listener.Addr().String()
}
func newTestClient(t *testing.T, serverAddress string) pb.ClientServiceClient {
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewClientServiceClient(conn)
}
