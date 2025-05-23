package usecase

import (
	"errors"
	"hexagonal-payment-api/internal/domain"
	"hexagonal-payment-api/internal/port"
)

type PaymentUseCase struct {
	repo port.PaymentService
}

func NewPaymentUseCase(repo port.PaymentService) *PaymentUseCase {
	return &PaymentUseCase{repo}
}

func (p *PaymentUseCase) CreatePayment(payment domain.Payment) error {
	if payment.Amount <= 0 {
		return errors.New("Invalid amount")
	}
	return p.repo.CreatePayment(payment)
}

func (p *PaymentUseCase) GetPaymentByUserID(userID string) ([]domain.Payment, error) {
	return p.repo.GetPaymentByUserID(userID)
}
