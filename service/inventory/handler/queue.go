package handler

import (
	"github.com/itzaddddd/game-shop/config"
	inventoryUsecase "github.com/itzaddddd/game-shop/service/inventory/usecase"
)

type (
	InventoryQueueHandlerService interface {
	}

	inventoryQueueHandler struct {
		cfg              *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryQueueHandler(cfg *config.Config, inventoryUsecase inventoryUsecase.InventoryUsecaseService) InventoryQueueHandlerService {
	return &inventoryQueueHandler{
		cfg:              cfg,
		inventoryUsecase: inventoryUsecase,
	}
}
