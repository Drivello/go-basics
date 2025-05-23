package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"payment-queue-api/dispatcher"
	"payment-queue-api/model"
	"payment-queue-api/worker"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	paymentInput := make(chan model.Payment)

	d := dispatcher.NewDispatcher()
	d.Start(ctx, paymentInput)

	go worker.StartWorker(ctx, "CardWorker", d.CardChan)
	go worker.StartWorker(ctx, "TransferWorker", d.TransferChan)
	go worker.StartWorker(ctx, "DebitWorker", d.DebitChan)

	// Simulate payments
	go func() {
		for i := 1; i <= 10; i++ {
			channel := model.Card
			switch i % 3 {
			case 0:
				channel = model.Card
			case 1:
				channel = model.Transfer
			case 2:
				channel = model.Debit
			}
			paymentInput <- model.Payment{
				ID:      i,
				Amount:  float64(i) * 100,
				Channel: channel,
			}
			time.Sleep(500 * time.Millisecond)
		}
		// simulate end of input
		close(paymentInput)
	}()

	<-sigChan
	fmt.Println("Shutting down system...")
	cancel()

	time.Sleep(3 * time.Second)
}
