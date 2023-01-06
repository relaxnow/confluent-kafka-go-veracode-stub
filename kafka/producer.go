package kafka

import "errors"

// Producer implements a High-level Apache Kafka Producer instance
type Producer struct {
	events         chan Event
	produceChannel chan *Message

	// Terminates the poller() goroutine
	pollerTermChan chan bool
}

func NewProducer(c *ConfigMap) (Producer, error) {
	return Producer{}, nil
}

// GetFatalError returns an Error object if the client instance has raised a fatal error, else nil.
func (p *Producer) GetFatalError() error {
	return getFatalError(p)
}

// ProduceChannel returns the produce *Message channel (write)
func (p *Producer) ProduceChannel() chan *Message {
	return p.produceChannel
}

func (p *Producer) Close() {
}

func (p *Producer) Produce(msg *Message, deliveryChan chan Event) error {
	return errors.New("Not implemented")
}

func (p *Producer) Flush(timeout int) int { return 0 }

// Events returns the Events channel (read)
func (p *Producer) Events() chan Event {
	return p.events
}
