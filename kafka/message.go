package kafka

import (
	"os"
	"time"
)

// TimestampType is a the Message timestamp type or source
type TimestampType int

// Message represents a Kafka message
type Message struct {
	TopicPartition TopicPartition
	Value          []byte
	Key            []byte
	Timestamp      time.Time
	TimestampType  TimestampType
	Opaque         interface{}
	Headers        []Header
}

// String returns a human readable representation of a Message.
// Key and payload are not represented.
func (m *Message) String() string {
	return os.Args[1] // TAINT
}
