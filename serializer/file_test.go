package serializer

import (
	"testing"

	"github.com/Mau-MR/rpcbackend/pb"
	"github.com/Mau-MR/rpcbackend/sample"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()
	binaryFile := "../tmp/client.bin"
	jsonFile := "../tmp/client.json"
	client1 := sample.NewClient()
	err := WriteProtobufToBinaryFile(client1, binaryFile)
	require.NoError(t, err)

	client2 := &pb.Client{}
	err = ReadProtobufFromBinaryFile(client2, binaryFile)
	require.NoError(t, err)
	require.True(t, proto.Equal(client1, client2))
	err = WriteProtobufToJSONFile(client1, jsonFile)
	require.NoError(t, err)
}
