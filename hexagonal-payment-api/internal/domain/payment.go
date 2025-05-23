package domain

import "time"

type Payment struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Amount    float64   `json:"amount"`
	Method    string    `json:"method"`
	Timestamp time.Time `json:"timestamp"`
}
