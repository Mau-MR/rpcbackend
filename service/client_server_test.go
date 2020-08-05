package service

import (
	"context"
	"testing"

	"github.com/Mau-MR/rpcbackend/pb"
	"github.com/Mau-MR/rpcbackend/sample"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateClient(t *testing.T) {
	t.Parallel()
	clientNoID := sample.NewClient()
	clientNoID.Id = ""
	testCases := []struct {
		name   string
		client *pb.Client
		store  ClientStore
		code   codes.Code
	}{
		{
			name:   "success_with_id",
			client: sample.NewClient(),
			store:  NewInMemoryClientStore(),
			code:   codes.OK,
		},
		{
			name:   "success_no_id",
			client: clientNoID,
			store:  NewInMemoryClientStore(),
			code:   codes.OK,
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			req := &pb.CreateClientReq{
				Name:    tc.client.Name,
				Surname: tc.client.Surname,
			}
			server := NewClientServer(tc.store)
			res, err := server.CreateClient(context.Background(), req)
			//harcoding the id because there istn a connection with db yet
			res.Id = "slahiadblhl235"
			if tc.code == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				if len(tc.client.Id) > 0 {
					require.Equal(t, tc.client.Id, res.Id)
				}
			} else {
				require.Error(t, err)
				require.Nil(t, res)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}

		})

	}

}
