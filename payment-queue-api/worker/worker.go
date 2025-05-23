package worker

import (
	"context"
	"fmt"
	"payment-queue-api/model"
	"time"
)

func StartWorker(ctx context.Context, name string, jobs <-chan model.Payment) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %s shutting down... \n", name)
		case payment, ok := <-jobs:
			if !ok {
				fmt.Printf("Channel for worker %s closed\n", name)
				return
			}
			fmt.Printf("Worker %s processing payment ID %d (%.2f)\n", name, payment.ID, payment.Amount)
			time.Sleep(1 * time.Second)
			fmt.Printf("Worker %s finished payment ID %d\n", name, payment.ID)
		}
	}
}
