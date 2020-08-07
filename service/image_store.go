package service

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
)

type ImageStore interface {
	Save(clientID string, imageType string, imageData bytes.Buffer) (string, error)
}

type DisckImageStore struct {
	mutex       sync.RWMutex
	imageFolder string
	images      map[string]*ImageInfo
}
type ImageInfo struct {
	ClientID string
	Type     string
	Path     string
}

func NewDiskImageStore(imageFolder string) *DisckImageStore {
	return &DisckImageStore{
		imageFolder: imageFolder,
		images:      make(map[string]*ImageInfo),
	}
}
func (store *DisckImageStore) Save(clientID string, imageType string, imageData bytes.Buffer) (string, error) {
	imageId, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("cannot generate imageId: %w", err)
	}
	imagePath := fmt.Sprintf("%s/%s%s", store.imageFolder, imageId, imageType)
	file, err := os.Create(imagePath)
	if err != nil {
		return "", fmt.Errorf("cannot create image file: %w", err)
	}
	_, err = imageData.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write image to file: %w", err)
	}
	store.mutex.Lock()
	defer store.mutex.Unlock()
	store.images[imageId.String()] = &ImageInfo{
		ClientID: imageId.String(),
		Type:     imageType,
		Path:     imagePath,
	}
	return imageId.String(), nil

}
