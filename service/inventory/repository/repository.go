package repository

import "go.mongodb.org/mongo-driver/mongo"

type (
	InventoryRepositoryService interface{}

	inventoryRepository struct {
		db *mongo.Client
	}
)

func NewInventoryRepository(db *mongo.Client) InventoryRepositoryService {
	return &inventoryRepository{db}
}
