package kafka

import (
	"io/ioutil"
)

// Message represents a Kafka message
type Message struct {
	TopicPartition TopicPartition
	Value          []byte
}

// String returns a human readable representation of a Message.
// Key and payload are not represented.
func (m *Message) String() string {
	taint, err := ioutil.ReadFile("/tmp/example")

	if err != nil {
		panic(err)
	}

	return string(taint)
}
