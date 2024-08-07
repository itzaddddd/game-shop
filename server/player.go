package server

import (
	"log"

	"github.com/itzaddddd/game-shop/pkg/grpccon"
	"github.com/itzaddddd/game-shop/service/player/handler"
	playerPb "github.com/itzaddddd/game-shop/service/player/proto"
	"github.com/itzaddddd/game-shop/service/player/repository"
	"github.com/itzaddddd/game-shop/service/player/usecase"
)

func (s *server) playerService() {
	repo := repository.NewPlayerRepository(s.db)
	usecase := usecase.NewPlayerUsecase(repo)
	httpHandler := handler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := handler.NewPlayerGrpcHandler(usecase)
	queueHandler := handler.NewPlayerQueueHandler(s.cfg, usecase)

	_ = grpcHandler
	_ = queueHandler

	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.PlayerUrl)
		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Player gPRC server listening on %s\n", s.cfg.Grpc.PlayerUrl)

		grpcServer.Serve(lis)

	}()

	playerRoute := s.app.Group("/player_v1")

	playerRoute.GET("", s.healthCheckService)
	playerRoute.POST("/player/register", httpHandler.CreatePlayer)
	playerRoute.GET("/player/:player_id", httpHandler.FindOnePlayerProfile)
	playerRoute.POST("/player/add-money", httpHandler.AddPlayerMoney)
	playerRoute.GET("/player/saving-account/my-account", httpHandler.GetPlayerSavingAccount)
}
