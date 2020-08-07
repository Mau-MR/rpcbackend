package service

import (
	"context"
	"errors"
	"fmt"
	"log"
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
	Find(id string) (*pb.Client, error)
	Search(ctx context.Context, filter *pb.ClientFilter, found func(client *pb.Client) error) error
	//	Db(client *pb.Client) (string, error)
}
type InMemoryClientStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Client
}
type DBClientStore struct {
}

func (store *InMemoryClientStore) Search(ctx context.Context, filter *pb.ClientFilter, found func(client *pb.Client) error) error {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	for _, client := range store.data {
		if ctx.Err() == context.Canceled || ctx.Err() == context.DeadlineExceeded {
			log.Printf("operation cancelled")
			return errors.New("context is cancelled")
		}
		if isQualified(filter, client) {
			other, err := deepCopy(client)
			if err != nil {
				return err
			}
			err = found(other)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func isQualified(filter *pb.ClientFilter, client *pb.Client) bool {
	if client.GetName() == filter.Name && client.GetSurname() == filter.Surname || client.GetPhone() == filter.Phone {
		return true
	}
	return false
}

//returns a new client store
func NewInMemoryClientStore() *InMemoryClientStore {
	return &InMemoryClientStore{
		data: make(map[string]*pb.Client),
	}
}
func (store *InMemoryClientStore) Find(id string) (*pb.Client, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()
	client := store.data[id]
	if client == nil {
		return nil, nil
	}
	return deepCopy(client)

}
func deepCopy(client *pb.Client) (*pb.Client, error) {
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
	if store.data[client.Id] != nil {
		return ErrAlreadyExists
	}
	other, err := deepCopy(client)
	if err != nil {
		return err
	}
	store.data[other.Id] = other
	return nil
}
