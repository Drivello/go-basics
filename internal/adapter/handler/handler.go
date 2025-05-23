package handler

import (
	"encoding/json"
	"hexagonal-payment-api/internal/domain"
	"hexagonal-payment-api/internal/usecase"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func NewRouter(usecase *usecase.PaymentUseCase) http.Handler {
	r := chi.NewRouter()

	r.Post("/payments", func(w http.ResponseWriter, r *http.Request) {
		var payment domain.Payment

		if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
			http.Error(w, "invalid payload", http.StatusBadRequest)
			return
		}

		payment.ID = uuid.New().String()
		payment.Timestamp = time.Now()

		if err := usecase.CreatePayment(payment); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	r.Get(("/payments/{user_id}"), func(w http.ResponseWriter, r *http.Request) {

		userID := chi.URLParam(r, "user_id")
		payments, err := usecase.GetPaymentByUserID(userID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(payments)

	})

	return r
}
