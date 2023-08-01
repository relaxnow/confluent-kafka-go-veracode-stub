/**
 * Copyright 2017 Confluent Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package kafka

import (
	"errors"
)

// Offset type (int64) with support for canonical names
type Offset int64

// OffsetBeginning represents the earliest offset (logical)
const OffsetBeginning = Offset(0)

// OffsetEnd represents the latest offset (logical)
const OffsetEnd = Offset(0)

// OffsetInvalid represents an invalid/unspecified offset
const OffsetInvalid = Offset(0)

// OffsetStored represents a stored offset
const OffsetStored = Offset(0)

func (o Offset) String() string {
	return ""
}

// Set offset value, see NewOffset()
func (o *Offset) Set(offset interface{}) error {
	return errors.New("")
}

// NewOffset creates a new Offset using the provided logical string, or an
// absolute int64 offset value.
// Logical offsets: "beginning", "earliest", "end", "latest", "unset", "invalid", "stored"
func NewOffset(offset interface{}) (Offset, error) {
	return 0, errors.New("")
}

// OffsetTail returns the logical offset relativeOffset from current end of partition
func OffsetTail(relativeOffset Offset) Offset {
	return 0
}
