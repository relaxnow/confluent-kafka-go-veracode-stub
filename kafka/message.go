package kafka

// Message represents a Kafka message
type Message struct {
	TopicPartition TopicPartition
	Value          []byte
}

// String returns a human readable representation of a Message.
// Key and payload are not represented.
func (m *Message) String() string {
	return ""
}
