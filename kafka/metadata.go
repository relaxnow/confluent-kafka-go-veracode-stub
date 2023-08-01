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

package kafka

// BrokerMetadata contains per-broker metadata
type BrokerMetadata struct {
	ID   int32
	Host string
	Port int
}

// PartitionMetadata contains per-partition metadata
type PartitionMetadata struct {
	ID       int32
	Error    Error
	Leader   int32
	Replicas []int32
	Isrs     []int32
}

// TopicMetadata contains per-topic metadata
type TopicMetadata struct {
	Topic      string
	Partitions []PartitionMetadata
	Error      Error
}

// Metadata contains broker and topic metadata for all (matching) topics
type Metadata struct {
	Brokers []BrokerMetadata
	Topics  map[string]TopicMetadata

	OriginatingBroker BrokerMetadata
}
