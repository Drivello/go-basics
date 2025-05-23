package main

import (
	"fmt"
	"time"

	"debt-scheduler/schedule"
)

func main() {
	req := schedule.ScheduleRequest{
		TotalAmount:  1000.00,
		Installments: 3,
		StartDate:    time.Date(2025, time.May, 22, 0, 0, 0, 0, time.Local),
	}

	payments := schedule.GenerateSchedule(req)

	for i, p := range payments {
		fmt.Printf("Cuota %d | Fecha: %s | Importe: $%.2f\n", i+1, p.DueDate.Format("02/01/2006"), p.Amount)
	}
}
