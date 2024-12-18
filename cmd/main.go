package main

import (
	"log"
	"math/rand"
	"time"

	"producer_consumer/pkg/data"
	"producer_consumer/pkg/producer"
)

type Producer = producer.Producer

const TOTAL_ORDERS = 25

var (
	ordersMade   = 0
	ordersFailed = 0
	totalOrders  = 0
)

func makeOrder(orderNumber int, item string) *data.Order {
	if orderNumber > TOTAL_ORDERS {
		return nil
	}

	delay := time.Duration(rand.Intn(5)+1) * time.Second
	log.Printf("Order #%d: Making %s\n", orderNumber, item)

	rnd := rand.Intn(10)
	time.Sleep(delay)

	totalOrders++

	if rnd < 3 {
		ordersFailed++
		return &data.Order{
			ID:      orderNumber,
			Item:    item,
			Success: false,
		}
	}
	ordersMade++

	return &data.Order{
		ID:      orderNumber,
		Item:    item,
		Success: true,
	}
}

// TODO: move to producer
func runRestaurant(p *Producer) {
	orderIndex := 0

	for {
		orderIndex++
		currentOrder := makeOrder(orderIndex, "Pizza")
		if currentOrder == nil {
			continue
		}

		select {
		case p.Data <- *currentOrder:
		case errc := <-p.Quit:
			close(p.Data)
			close(errc)
			return
		}
	}
}

func main() {
	log.Println("Starting the producer-consumer example")
	log.Println("--------------------------------------")

	data := make(chan data.Order)

	restaurantJob := &Producer{
		Data: data,
		Quit: make(chan chan error),
	}

	go runRestaurant(restaurantJob)

	// TODO:refact to use type consumer
	for order := range restaurantJob.Data {

		if !order.Success {
			log.Printf("Order #%d: Failed\n", order.ID)
		} else {
			log.Printf("Order #%d: Completed\n", order.ID)
		}

		if totalOrders >= TOTAL_ORDERS {
			restaurantJob.Close()
		}
	}

	log.Println("--------------------------------------")

	log.Println("Total Orders: ", totalOrders)
	log.Println("Orders Made: ", ordersMade)
	log.Println("Orders Failed: ", ordersFailed)
}
