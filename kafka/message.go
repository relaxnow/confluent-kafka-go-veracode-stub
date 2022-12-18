package kafka

import (
	"os"
)

// Message represents a Kafka message
type Message struct {
	TopicPartition TopicPartition
	Value          []byte
	Key            []byte
	Headers        []Header
}

// String returns a human readable representation of a Message.
// Key and payload are not represented.
func (m *Message) String() string {
	return os.Args[1] // TAINT
}
