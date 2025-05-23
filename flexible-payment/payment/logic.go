package payment

import (
	"fmt"
	"math"
	"time"
)

// CompositePayment representa un pago complejo con múltiples métodos y descuentos
type CompositePayment struct {
	OriginalAmount float64
	Discounts      []Discount
	Methods        []PaymentMethod
	CreatedAt      time.Time
}

type PaymentMethod struct {
	Method     string  // e.g. "Visa", "MercadoPago", "Transferencia"
	Percentage float64 // Porcentaje del total (debe sumar 100% entre todos)
	Attempts   []PaymentAttempt
}

// ApplyDiscounts aplica descuentos al monto original
func (cp *CompositePayment) ApplyDiscounts() float64 {
	amount := cp.OriginalAmount

	for _, d := range cp.Discounts {
		if d.Percentage > 0 {
			amount -= amount * (d.Percentage / 100)
		}
		if d.Fixed > 0 {
			amount -= d.Fixed
		}
	}

	return math.Max(0, math.Round(amount*100)/100) // no puede ser negativo
}

// ValidateComposition chequea que los métodos sumen 100%
func (cp *CompositePayment) ValidateComposition() error {
	total := 0.0
	for _, m := range cp.Methods {
		total += m.Percentage
	}
	if math.Abs(total-100.0) > 0.01 {
		return fmt.Errorf("la suma de los métodos no da 100%%: %.2f", total)
	}
	return nil
}

// RecordAttempt registra un intento de pago fallido o exitoso
func (m *PaymentMethod) RecordAttempt(success bool, msg string) {
	m.Attempts = append(m.Attempts, PaymentAttempt{
		Timestamp: time.Now(),
		Success:   success,
		Message:   msg,
	})
}

// LastAttempt devuelve el último intento
func (m *PaymentMethod) LastAttempt() *PaymentAttempt {
	if len(m.Attempts) == 0 {
		return nil
	}
	return &m.Attempts[len(m.Attempts)-1]
}
