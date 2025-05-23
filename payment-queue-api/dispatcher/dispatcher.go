package dispatcher

import (
	"context"
	"payment-queue-api/model"
)

type Dispatcher struct {
	CardChan     chan model.Payment
	TransferChan chan model.Payment
	DebitChan    chan model.Payment
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		CardChan:     make(chan model.Payment),
		TransferChan: make(chan model.Payment),
		DebitChan:    make(chan model.Payment),
	}
}

func (d *Dispatcher) Start(ctx context.Context, input <-chan model.Payment) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(d.CardChan)
				close(d.TransferChan)
				close(d.DebitChan)
				return
			case payment, ok := <-input:
				if !ok {
					return
				}
				switch payment.Channel {
				case model.Card:
					d.CardChan <- payment
				case model.Transfer:
					d.TransferChan <- payment
				case model.Debit:
					d.DebitChan <- payment
				}
			}
		}
	}()
}
