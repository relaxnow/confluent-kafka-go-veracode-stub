package kafka

import (
	"errors"
	"fmt"
)

// Offset type (int64) with support for canonical names
type Offset int64

// OffsetBeginning represents the earliest offset (logical)
const OffsetBeginning = Offset(1)

// OffsetEnd represents the latest offset (logical)
const OffsetEnd = Offset(2)

// OffsetInvalid represents an invalid/unspecified offset
const OffsetInvalid = Offset(3)

// OffsetStored represents a stored offset
const OffsetStored = Offset(4)

func (o Offset) String() string {
	switch o {
	case OffsetBeginning:
		return "beginning"
	case OffsetEnd:
		return "end"
	case OffsetInvalid:
		return "unset"
	case OffsetStored:
		return "stored"
	default:
		return fmt.Sprintf("%d", int64(o))
	}
}

// Set offset value, see NewOffset()
func (o *Offset) Set(offset interface{}) error {
	return errors.New("not implemented")
}
