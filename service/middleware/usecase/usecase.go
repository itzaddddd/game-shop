package usecase

import middlewareRepository "github.com/itzaddddd/game-shop/service/middleware/repository"

type (
	MiddlewareUsecaseService interface{}

	middlewareUsecase struct {
		middlewareRepository middlewareRepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewareRepository.MiddlewareRepositoryService) MiddlewareUsecaseService {
	return &middlewareUsecase{middlewareRepository}
}
