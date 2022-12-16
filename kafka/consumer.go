package kafka

import (
	"errors"
	"time"
)

// RebalanceCb provides a per-Subscribe*() rebalance event callback.
// The passed Event will be either AssignedPartitions or RevokedPartitions
type RebalanceCb func(*Consumer, Event) error

// Consumer implements a High-level Apache Kafka Consumer instance
type Consumer struct {
}

func NewConsumer(conf *ConfigMap) (*Consumer, error) {
	return &Consumer{}, errors.New("")
}

func (c *Consumer) SubscribeTopics(topics []string, rebalanceCb RebalanceCb) (err error) {
	return errors.New("")
}

func (c *Consumer) ReadMessage(timeout time.Duration) (*Message, error) {
	return &Message{}, errors.New("")
}

func (c *Consumer) Close() (err error) {
	return errors.New("")
}
