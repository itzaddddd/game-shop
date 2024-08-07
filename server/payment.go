package server

import (
	"github.com/itzaddddd/game-shop/service/payment/handler"
	"github.com/itzaddddd/game-shop/service/payment/repository"
	"github.com/itzaddddd/game-shop/service/payment/usecase"
)

func (s *server) paymentService() {
	repo := repository.NewPaymentRepository(s.db)
	usecase := usecase.NewPaymentUsecase(repo)
	httpHandler := handler.NewPaymentHttpHandler(s.cfg, usecase)

	_ = httpHandler

	paymentRoute := s.app.Group("/payment_v1")

	paymentRoute.GET("", s.healthCheckService)
}
