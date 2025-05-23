package schedule

import (
	"math"
)

// GenerateSchedule returns a slice of monthly payments evenly splitting the total
func GenerateSchedule(req ScheduleRequest) []Payment {
	payments := make([]Payment, req.Installments)

	// Calculate base amount per installment (rounded to 2 decimals)
	base := math.Floor((req.TotalAmount/float64(req.Installments))*100) / 100
	totalAssigned := base * float64(req.Installments)

	// Calculate leftover due to rounding
	remainder := math.Round((req.TotalAmount-totalAssigned)*100) / 100

	for i := 0; i < req.Installments; i++ {
		amount := base
		// Distribute remaining cents to the first installments
		if remainder > 0 {
			amount += 0.01
			remainder -= 0.01
		}

		dueDate := req.StartDate.AddDate(0, i, 0) // Add i months
		payments[i] = Payment{
			DueDate: dueDate,
			Amount:  amount,
		}
	}

	return payments
}
