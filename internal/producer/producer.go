package producer

import (
	"math/rand"
	"producer_consumer/internal/order"
	"time"
)

type Producer struct {
	Data chan order.Order
	Quit chan chan error
}

func (p *Producer) Start() {
	orderNumber := 0
	for {
		orderNumber++
		select {
		case <-p.Quit:
			return
		case p.Data <- order.Order{ID: orderNumber, Success: rand.Intn(2) == 0}:
			// TODO: user Logger
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}
}
