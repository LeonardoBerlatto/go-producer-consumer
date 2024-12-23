package consumer

import (
	"log"
	"producer_consumer/internal/order"
)

type Consumer struct {
	Data chan order.Order
}

func (c *Consumer) Start() {
	for currentOrder := range c.Data {

		// TODO: user Logger
		if !currentOrder.Success {
			log.Printf("Order #%d: Failed\n", currentOrder.ID)
		} else {
			log.Printf("Order #%d: Completed\n", currentOrder.ID)
		}
	}

}
