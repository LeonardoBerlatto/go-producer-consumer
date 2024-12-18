package producer

import (
	"producer_consumer/pkg/data"
)

type Producer struct {
	Data chan data.Order
	Quit chan chan error
}

func (p *Producer) Close() error {
	errc := make(chan error)
	p.Quit <- errc
	return <-errc
}
