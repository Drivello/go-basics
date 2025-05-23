package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Payment struct {
	Source string
	ID     int
	Amount float64
}

func paymentSource(ctx context.Context, name string, out chan<- Payment) {
	delay := time.Duration(rand.Intn(1500)) * time.Millisecond
	select {
	case <-time.After(delay):
		select {
		case out <- Payment{Source: name, ID: rand.Intn(1000), Amount: rand.Float64() * 1000}:
		case <-ctx.Done():
		}
	case <-ctx.Done():
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	output := make(chan Payment)

	sources := []string{"Visa", "Mastercard", "MercadoPago", "Transferencia"}

	for _, source := range sources {
		go paymentSource(ctx, source, output)
	}

	select {
	case p := <-output:
		fmt.Printf("Primer pago recibido de %s: ID=%d, Monto=%.2f\n", p.Source, p.ID, p.Amount)
	case <-ctx.Done():
		fmt.Println("Timeout: ninguna fuente respondió a tiempo")
	}
}
