package handler

import (
	"context"

	itemPb "github.com/itzaddddd/game-shop/service/item/proto"
	itemUsecase "github.com/itzaddddd/game-shop/service/item/usecase"
)

type (
	itemGrpcHandler struct {
		itemPb.UnimplementedItemGrpcServiceServer
		itemUsecase itemUsecase.ItemUsecaseService
	}
)

func NewItemGrpcHandler(itemUsecase itemUsecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{
		itemUsecase: itemUsecase,
	}
}

func (h *itemGrpcHandler) FindItemsInIds(context.Context, *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error) {
	return nil, nil
}
