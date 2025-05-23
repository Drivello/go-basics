package main

import (
	"hexagonal-payment-api/internal/adapter/handler"
	"hexagonal-payment-api/internal/adapter/repository"
	"hexagonal-payment-api/internal/usecase"
	"hexagonal-payment-api/pkg/middleware"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	inMemoryPaymentRepository := repository.NewInMemoryPaymentRepository()
	paymentUseCases := usecase.NewPaymentUseCase(inMemoryPaymentRepository)
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Mount("/", handler.NewRouter(paymentUseCases))

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
