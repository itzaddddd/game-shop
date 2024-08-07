package server

import (
	"log"

	"github.com/itzaddddd/game-shop/pkg/grpccon"
	"github.com/itzaddddd/game-shop/service/item/handler"
	itemPb "github.com/itzaddddd/game-shop/service/item/proto"
	"github.com/itzaddddd/game-shop/service/item/repository"
	"github.com/itzaddddd/game-shop/service/item/usecase"
)

func (s *server) itemService() {
	repo := repository.NewItemRepository(s.db)
	usecase := usecase.NewItemUsecase(repo)
	httpHandler := handler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := handler.NewItemGrpcHandler(usecase)

	_ = httpHandler

	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ItemUrl)
		itemPb.RegisterItemGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Item gPRC server listening on %s\n", s.cfg.Grpc.ItemUrl)

		grpcServer.Serve(lis)

	}()

	itemRoute := s.app.Group("/item_v1")

	itemRoute.GET("", s.healthCheckService)
}
