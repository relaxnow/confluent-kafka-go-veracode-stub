package kafka

// ErrorCode is the integer representation of local and broker error codes
type ErrorCode int

const ErrAllBrokersDown ErrorCode = 1

// String returns a human readable representation of an error code
func (c ErrorCode) String() string {
	return string(c)
}
