package handler

import (
	"github.com/itzaddddd/game-shop/config"
	playerUsecase "github.com/itzaddddd/game-shop/service/player/usecase"
)

type (
	PlayerQueueHandlerService interface {
	}

	playerQueueHandler struct {
		cfg           *config.Config
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerQueueHandler(cfg *config.Config, playerUsecase playerUsecase.PlayerUsecaseService) PlayerQueueHandlerService {
	return &playerQueueHandler{
		cfg:           cfg,
		playerUsecase: playerUsecase,
	}
}
