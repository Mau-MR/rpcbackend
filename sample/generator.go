package sample

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Mau-MR/rpcbackend/pb"
	"github.com/google/uuid"
)

//returns a newclient for testing
func init() {
	rand.Seed(time.Now().UnixNano())
}
func NewClient() *pb.Client {

	names := [4]string{"Maricio", "Marco", "Raul", "Franciso"}
	surnames := [4]string{"Alcantara", "Soriano", "Opul", "Fritz"}
	id, err := uuid.NewRandom()
	if err != nil {
		log.Fatal("cannot generate new uuid")
	}
	phones := 1000000000 + rand.Int()%(9999999999-1000000000)

	client := &pb.Client{
		Id:      id.String(),
		Name:    names[rand.Intn(4)],
		Surname: surnames[rand.Intn(4)],
		Phone:   strconv.Itoa(phones),
		App:     rand.Intn(2) == 1,
	}
	return client

}
