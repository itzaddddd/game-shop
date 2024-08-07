package handler

import (
	"context"

	authPb "github.com/itzaddddd/game-shop/service/auth/proto"
	authUsecase "github.com/itzaddddd/game-shop/service/auth/usecase"
)

type (
	authGrpcHndler struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authUsecase.AuthUsecaseService) *authGrpcHndler {
	return &authGrpcHndler{
		authUsecase: authUsecase,
	}
}

func (h *authGrpcHndler) AccessTokenSearch(ctx context.Context, req *authPb.AccessTokenSearchReq) (*authPb.AccessTokenSearchRes, error) {
	return nil, nil
}

func (h *authGrpcHndler) RolesCount(ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return nil, nil
}
