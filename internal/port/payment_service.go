package port

import "hexagonal-payment-api/internal/domain"

type PaymentService interface {
	CreatePayment(payment domain.Payment) error
	GetPaymentByUserID(userID string) ([]domain.Payment, error)
}
