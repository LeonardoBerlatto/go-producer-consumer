package consumer

import (
	"go.uber.org/zap"
	"producer_consumer/internal/order"
	logger "producer_consumer/pkg/log"
)

var log = logger.GetLogger()

type Consumer struct {
	Data chan order.Order
}

func (c *Consumer) Start() {
	for currentOrder := range c.Data {

		if !currentOrder.Success {
			log.Warn("Order failed", zap.Int("orderNumber", currentOrder.ID))
		} else {
			log.Info("Order completed", zap.Int("orderNumber", currentOrder.ID))
		}
	}

}
