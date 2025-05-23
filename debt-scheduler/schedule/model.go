package schedule

import "time"

// ScheduleRequest defines input for debt distribution
type ScheduleRequest struct {
	TotalAmount  float64   // e.g. 1000.00
	Installments int       // e.g. 3
	StartDate    time.Time // first due date
}

// Payment represents one monthly payment
type Payment struct {
	DueDate time.Time
	Amount  float64
}
