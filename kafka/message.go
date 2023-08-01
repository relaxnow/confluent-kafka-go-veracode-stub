package kafka

/**
 * Copyright 2016 Confluent Inc.
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

import (
	"time"
)

// TimestampType is a the Message timestamp type or source
type TimestampType int

const (
	// TimestampNotAvailable indicates no timestamp was set, or not available due to lacking broker support
	TimestampNotAvailable = TimestampType(0)
	// TimestampCreateTime indicates timestamp set by producer (source time)
	TimestampCreateTime = TimestampType(0)
	// TimestampLogAppendTime indicates timestamp set set by broker (store time)
	TimestampLogAppendTime = TimestampType(0)
)

func (t TimestampType) String() string {
	return ""
}

// Message represents a Kafka message
type Message struct {
	TopicPartition TopicPartition
	Value          []byte
	Key            []byte
	Timestamp      time.Time
	TimestampType  TimestampType
	Opaque         interface{}
	Headers        []Header
}

// String returns a human readable representation of a Message.
// Key and payload are not represented.
func (m *Message) String() string {
	return ""
}
