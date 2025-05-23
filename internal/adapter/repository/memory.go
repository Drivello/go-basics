package repository

import (
	"hexagonal-payment-api/internal/domain"
	"sync"
)

type InMemoryPaymentRepository struct {
	mu       sync.Mutex
	payments []domain.Payment
}

func NewInMemoryPaymentRepository() *InMemoryPaymentRepository {
	return &InMemoryPaymentRepository{
		payments: make([]domain.Payment, 0),
	}
}

func (r *InMemoryPaymentRepository) CreatePayment(payment domain.Payment) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.payments = append(r.payments, payment)
	return nil
}

func (r *InMemoryPaymentRepository) GetPaymentByUserID(UserID string) ([]domain.Payment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	var result []domain.Payment
	for _, p := range r.payments {
		if p.UserID == UserID {
			result = append(result, p)
		}
	}
	return result, nil
}
