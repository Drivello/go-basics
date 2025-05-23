package main

import (
	"fmt"
	"time"

	"flexible-payment/payment"
)

func main() {
	p := payment.CompositePayment{
		OriginalAmount: 1000,
		Discounts: []payment.Discount{
			{Description: "Descuento por promoción", Percentage: 10},
			{Description: "Cupón fijo", Fixed: 100},
		},
		Methods: []payment.PaymentMethod{
			{Method: "Visa", Percentage: 50},
			{Method: "MercadoPago", Percentage: 50},
		},
		CreatedAt: time.Now(),
	}

	// Validación
	if err := p.ValidateComposition(); err != nil {
		panic(err)
	}

	final := p.ApplyDiscounts()
	fmt.Printf("Monto original: $%.2f\n", p.OriginalAmount)
	fmt.Printf("Monto con descuentos: $%.2f\n", final)

	// Simular reintentos
	p.Methods[0].RecordAttempt(false, "Tarjeta rechazada por límite")
	p.Methods[0].RecordAttempt(true, "Aprobado en segundo intento")

	fmt.Println("Historial de Visa:")
	for _, attempt := range p.Methods[0].Attempts {
		fmt.Printf("  [%s] %s - %v\n", attempt.Timestamp.Format("15:04:05"), attempt.Message, attempt.Success)
	}
}
