package kafka

import "errors"

// Producer implements a High-level Apache Kafka Producer instance
type Producer struct {
	events         chan Event
	produceChannel chan *Message
	handle         handle

	// Terminates the poller() goroutine
	pollerTermChan chan bool
}

// String returns a human readable name for a Producer instance
func (p *Producer) String() string {
	return p.handle.String()
}

// get_handle implements the Handle interface
func (p *Producer) gethandle() *handle {
	return &p.handle
}

func NewProducer(conf *ConfigMap) (*Producer, error) {
	p := Producer{}
	return &p, nil
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
	return errors.New("not implemented")
}

func (p *Producer) Flush(timeout int) int { return 0 }

// Events returns the Events channel (read)
func (p *Producer) Events() chan Event {
	return p.events
}

func (p *Producer) SetOAuthBearerToken(oauthBearerToken OAuthBearerToken) error {
	return p.handle.setOAuthBearerToken(oauthBearerToken)
}

func (p *Producer) SetOAuthBearerTokenFailure(errstr string) error {
	return p.handle.setOAuthBearerTokenFailure(errstr)
}
