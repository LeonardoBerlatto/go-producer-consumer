package producer

import (
	"go.uber.org/zap"
	"math/rand"
	"producer_consumer/internal/order"
	"producer_consumer/pkg/log"
	"time"
)

var log = logger.GetLogger()

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
		case p.Data <- order.Order{ID: orderNumber, Success: generateOrderOutcome()}:
			log.Info("Order sent", zap.Int("orderNumber", orderNumber))
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}
}

func generateOrderOutcome() bool {
	return rand.Intn(2) == 0
}
