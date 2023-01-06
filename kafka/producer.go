package kafka

import "errors"

// Producer implements a High-level Apache Kafka Producer instance
type Producer struct {
	events chan Event
}

func NewProducer(c *ConfigMap) (Producer, error) {
	return Producer{}, nil
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
