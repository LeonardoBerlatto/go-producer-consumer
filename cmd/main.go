package main

import (
	"os"
	"os/signal"
	"producer_consumer/internal/consumer"
	"producer_consumer/internal/order"
	"producer_consumer/internal/producer"
	"producer_consumer/pkg/log"
	"syscall"
)

func main() {
	log := logger.GetLogger()
	log.Info("Starting the producer-consumer example")
	log.Info("-------------------------------------")

	buffer := make(chan order.Order)
	quit := make(chan chan error)

	producerJob := &producer.Producer{
		Data: buffer,
		Quit: make(chan chan error),
	}

	consumerJob := consumer.Consumer{
		Data: buffer,
	}

	go producerJob.Start()
	go consumerJob.Start()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	<-signalChan
	close(quit)
	close(buffer)

	log.Info("-------------------------------------")
	log.Info("Shutting down the producer-consumer example")
}
