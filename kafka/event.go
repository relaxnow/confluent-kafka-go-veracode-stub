package kafka

// Event generic interface
type Event interface {
	// String returns a human-readable representation of the event
	String() string
}
