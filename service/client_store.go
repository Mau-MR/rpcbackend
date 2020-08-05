package service

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Mau-MR/rpcbackend/pb"
	"github.com/jinzhu/copier"
)

//ErrAlreadyExists is returned when a record with the same id already exist on db or store
var ErrAlreadyExists = errors.New("record already exists")

//interface to store interfaces
type ClientStore interface {
	//Save laptop to store
	Save(client *pb.Client) error
	//TODO:Find the right structure
	//this could be changed for id
	Find(phone string) (*pb.Client, error)
	//	Db(client *pb.Client) (string, error)
}
type InMemoryClientStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Client
}
type DBClientStore struct {
}

//returns a new client store
func NewInMemoryClientStore() *InMemoryClientStore {
	return &InMemoryClientStore{
		data: make(map[string]*pb.Client),
	}
}
func (store *InMemoryClientStore) Find(phone string) (*pb.Client, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	client := store.data[phone]
	if client == nil {
		return nil, nil
	}
	//this just for not having problems with passing a pointer
	other := &pb.Client{}
	err := copier.Copy(other, client)
	if err != nil {
		return nil, fmt.Errorf("cannot copy client data: %w", err)
	}
	return other, nil

}
func (store *InMemoryClientStore) Save(client *pb.Client) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()
	if store.data[client.Phone] != nil {
		return ErrAlreadyExists
	}
	//TODO: find out if it is a better to use coppier
	store.data[client.Phone] = client
	return nil
}
