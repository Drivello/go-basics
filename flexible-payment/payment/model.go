package payment

import "time"

type PaymentAttempt struct {
	Timestamp time.Time
	Success   bool
	Message   string
}

type Discount struct {
	Description string
	Fixed       float64 // Monto fijo
	Percentage  float64 // % del total
}
