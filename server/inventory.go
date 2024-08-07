package server

import (
	"github.com/itzaddddd/game-shop/service/inventory/handler"
	"github.com/itzaddddd/game-shop/service/inventory/repository"
	"github.com/itzaddddd/game-shop/service/inventory/usecase"
)

func (s *server) inventoryService() {
	repo := repository.NewInventoryRepository(s.db)
	usecase := usecase.NewInventoryUsecase(repo)
	httpHandler := handler.NewInventoryHttpHandler(s.cfg, usecase)
	queueHandler := handler.NewInventoryQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = queueHandler

	inventoryRoute := s.app.Group("/inventory_v1")

	inventoryRoute.GET("", s.healthCheckService)
}
