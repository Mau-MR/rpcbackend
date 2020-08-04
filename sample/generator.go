package sample

import "github.com/Mau-MR/rpcbackend/pb"

//returns a newclient for testing
func NewClient() *pb.Client {
	client := &pb.Client{
		Name:    "Mauricio",
		Surname: "Merida",
		Phone:   "7773491106",
		App:     false,
	}
	return client

}
