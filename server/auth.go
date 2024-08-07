package server

import (
	"log"

	"github.com/itzaddddd/game-shop/pkg/grpccon"
	"github.com/itzaddddd/game-shop/service/auth/handler"
	authPb "github.com/itzaddddd/game-shop/service/auth/proto"
	"github.com/itzaddddd/game-shop/service/auth/repository"
	"github.com/itzaddddd/game-shop/service/auth/usecase"
)

func (s *server) authService() {
	repo := repository.NewAuthRepository(s.db)
	usecase := usecase.NewAuthUsecase(repo)
	httpHandler := handler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := handler.NewAuthGrpcHandler(usecase)

	_ = httpHandler

	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.AuthUrl)
		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Auth gPRC server listening on %s\n", s.cfg.Grpc.AuthUrl)

		grpcServer.Serve(lis)

	}()

	authRoute := s.app.Group("/auth_v1")

	authRoute.GET("", s.healthCheckService)
}
