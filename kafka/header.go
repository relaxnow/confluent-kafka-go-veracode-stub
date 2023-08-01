package kafka

/**
 * Copyright 2018 Confluent Inc.
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

// Header represents a single Kafka message header.
//
// Message headers are made up of a list of Header elements, retaining their original insert
// order and allowing for duplicate Keys.
//
// Key is a human readable string identifying the header.
// Value is the key's binary value, Kafka does not put any restrictions on the format of
// of the Value but it should be made relatively compact.
// The value may be a byte array, empty, or nil.
//
// NOTE: Message headers are not available on producer delivery report messages.
type Header struct {
	Key   string // Header name (utf-8 string)
	Value []byte // Header value (nil, empty, or binary)
}

// String returns the Header Key and data in a human representable possibly truncated form
// suitable for displaying to the user.
func (h Header) String() string {
	return ""
}
