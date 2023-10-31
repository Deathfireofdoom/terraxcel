package client

import (
	"fmt"

	"github.com/Deathfireofdoom/terraxcel/server/pkg/repository"
)

type TerraxcelClient struct {
	repository *repository.DocumentRepository
}

func NewTerraxcelClient() (*TerraxcelClient, error) {
	repository, err := repository.NewDocumentRepository()

	if err != nil {
		fmt.Printf("failed to create repository: %v", err)
		return nil, err
	}

	return &TerraxcelClient{repository: repository}, nil
}
