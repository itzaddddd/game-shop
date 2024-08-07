package usecase

import paymentRepository "github.com/itzaddddd/game-shop/service/payment/repository"

type (
	PaymentUsecaseService interface {
	}

	paymentUsecase struct {
		paymentRepository paymentRepository.PaymentRepositoryService
	}
)

func NewPaymentUsecase(paymentRepository paymentRepository.PaymentRepositoryService) PaymentUsecaseService {
	return &paymentUsecase{
		paymentRepository: paymentRepository,
	}
}
