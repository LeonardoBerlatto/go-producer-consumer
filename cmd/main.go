package main

import (
	"os"
	"os/signal"
	"producer_consumer/internal/consuming"
	"producer_consumer/internal/item"
	"producer_consumer/internal/producing"
	"producer_consumer/pkg/log"
	"syscall"
)

func main() {
	log := logger.GetLogger()
	log.Info("Starting the producing-consuming example")
	log.Info("-------------------------------------")

	buffer := make(chan item.Order)
	quit := make(chan chan error)

	producerJob := &producing.Producer{
		Data: buffer,
		Quit: make(chan chan error),
	}

	consumerJob := consuming.Consumer{
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
	log.Info("Shutting down the producing-consuming example")
}
