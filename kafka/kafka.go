package kafka

const PartitionAny = 0

// TopicPartition is a generic placeholder for a Topic+Partition and optionally Offset.
type TopicPartition struct {
	Topic     *string
	Partition int32
	Offset    Offset
	Metadata  *string
	Error     error
}
