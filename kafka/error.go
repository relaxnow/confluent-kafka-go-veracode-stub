package kafka

import "fmt"

// Error provides a Kafka-specific error container
type Error struct {
	code             ErrorCode
	str              string
	fatal            bool
	retriable        bool
	txnRequiresAbort bool
}

// String returns a human readable representation of an Error
func (e Error) String() string {
	var errstr string
	if len(e.str) > 0 {
		errstr = e.str
	} else {
		errstr = e.code.String()
	}

	if e.IsFatal() {
		return fmt.Sprintf("Fatal error: %s", errstr)
	}

	return errstr
}

func (e Error) Code() ErrorCode {
	return e.code
}

func (e Error) IsFatal() bool {
	return e.fatal
}

// Error returns a human readable representation of an Error
// Same as Error.String()
func (e Error) Error() string {
	return e.String()
}
