package usecase

import itemRepository "github.com/itzaddddd/game-shop/service/item/repository"

type (
	ItemUsecaseService interface {
	}

	itemUsecase struct {
		itemRepository itemRepository.ItemRepositoryService
	}
)

func NewItemUsecase(itemRepository itemRepository.ItemRepositoryService) ItemUsecaseService {
	return &itemUsecase{itemRepository}
}
