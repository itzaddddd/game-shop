package handler

import (
	"github.com/itzaddddd/game-shop/config"
	paymentUsecase "github.com/itzaddddd/game-shop/service/payment/usecase"
)

type (
	PaymentHttpHandlerService interface {
	}

	paymentHttpHandler struct {
		cfg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler(cfg *config.Config, paymentUsecase paymentUsecase.PaymentUsecaseService) PaymentHttpHandlerService {
	return &paymentHttpHandler{
		cfg:            cfg,
		paymentUsecase: paymentUsecase,
	}
}
