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

func (c *Consumer) Assign(partitions []TopicPartition) (err error) {
	return nil
}

// Assignment returns the current partition assignments
func (c *Consumer) Assignment() (partitions []TopicPartition, err error) {
	return nil, nil
}

// CommitOffsets commits the provided list of offsets
// This is a blocking call.
// Returns the committed offsets on success.
func (c *Consumer) CommitOffsets(offsets []TopicPartition) ([]TopicPartition, error) {
	return nil, nil
}

func (c *Consumer) Pause(partitions []TopicPartition) (err error) {
	return nil
}

func (c *Consumer) Poll(timeoutMs int) (event Event) {
	return nil
}

// Resume consumption for the provided list of partitions
func (c *Consumer) Resume(partitions []TopicPartition) (err error) {
	return nil
}

func (c *Consumer) StoreOffsets(offsets []TopicPartition) (storedOffsets []TopicPartition, err error) {
	return nil, nil
}

func (c *Consumer) Unassign() (err error) {
	return nil
}

// Logs returns the log channel if enabled, or nil otherwise.
func (c *Consumer) Logs() chan LogEvent {
	return nil
}
