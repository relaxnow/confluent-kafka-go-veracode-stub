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
	"errors"
	"time"
)

// RebalanceCb provides a per-Subscribe*() rebalance event callback.
// The passed Event will be either AssignedPartitions or RevokedPartitions
type RebalanceCb func(*Consumer, Event) error

// Consumer implements a High-level Apache Kafka Consumer instance
type Consumer struct {
	events             chan Event
	handle             handle
	eventsChanEnable   bool
	readerTermChan     chan bool
	rebalanceCb        RebalanceCb
	appReassigned      bool
	appRebalanceEnable bool // config setting
}

// Strings returns a human readable name for a Consumer instance
func (c *Consumer) String() string {
	return ""
}

// getHandle implements the Handle interface
func (c *Consumer) gethandle() *handle {
	return nil
}

// Subscribe to a single topic
// This replaces the current subscription
func (c *Consumer) Subscribe(topic string, rebalanceCb RebalanceCb) error {
	return errors.New("")
}

// SubscribeTopics subscribes to the provided list of topics.
// This replaces the current subscription.
func (c *Consumer) SubscribeTopics(topics []string, rebalanceCb RebalanceCb) (err error) {
	return errors.New("")
}

// Unsubscribe from the current subscription, if any.
func (c *Consumer) Unsubscribe() (err error) {
	return errors.New("")
}

// Assign an atomic set of partitions to consume.
// This replaces the current assignment.
func (c *Consumer) Assign(partitions []TopicPartition) (err error) {
	return errors.New("")
}

// Unassign the current set of partitions to consume.
func (c *Consumer) Unassign() (err error) {
	return errors.New("")
}

// commit offsets for specified offsets.
// If offsets is nil the currently assigned partitions' offsets are committed.
// This is a blocking call, caller will need to wrap in go-routine to
// get async or throw-away behaviour.
func (c *Consumer) commit(offsets []TopicPartition) (committedOffsets []TopicPartition, err error) {
	return nil, errors.New("")
}

// Commit offsets for currently assigned partitions
// This is a blocking call.
// Returns the committed offsets on success.
func (c *Consumer) Commit() ([]TopicPartition, error) {
	return nil, errors.New("")
}

// CommitMessage commits offset based on the provided message.
// This is a blocking call.
// Returns the committed offsets on success.
func (c *Consumer) CommitMessage(m *Message) ([]TopicPartition, error) {
	return nil, errors.New("")
}

// CommitOffsets commits the provided list of offsets
// This is a blocking call.
// Returns the committed offsets on success.
func (c *Consumer) CommitOffsets(offsets []TopicPartition) ([]TopicPartition, error) {
	return nil, errors.New("")
}

// StoreOffsets stores the provided list of offsets that will be committed
// to the offset store according to `auto.commit.interval.ms` or manual
// offset-less Commit().
//
// Returns the stored offsets on success. If at least one offset couldn't be stored,
// an error and a list of offsets is returned. Each offset can be checked for
// specific errors via its `.Error` member.
func (c *Consumer) StoreOffsets(offsets []TopicPartition) (storedOffsets []TopicPartition, err error) {
	return nil, errors.New("")
}

// Seek seeks the given topic partitions using the offset from the TopicPartition.
//
// If timeoutMs is not 0 the call will wait this long for the
// seek to be performed. If the timeout is reached the internal state
// will be unknown and this function returns ErrTimedOut.
// If timeoutMs is 0 it will initiate the seek but return
// immediately without any error reporting (e.g., async).
//
// Seek() may only be used for partitions already being consumed
// (through Assign() or implicitly through a self-rebalanced Subscribe()).
// To set the starting offset it is preferred to use Assign() and provide
// a starting offset for each partition.
//
// Returns an error on failure or nil otherwise.
func (c *Consumer) Seek(partition TopicPartition, timeoutMs int) error {
	return errors.New("")
}

// Poll the consumer for messages or events.
//
// # Will block for at most timeoutMs milliseconds
//
// The following callbacks may be triggered:
//
//	Subscribe()'s rebalanceCb
//
// Returns nil on timeout, else an Event
func (c *Consumer) Poll(timeoutMs int) (event Event) {
	return nil
}

// Events returns the Events channel (if enabled)
func (c *Consumer) Events() chan Event {
	return nil
}

// Logs returns the log channel if enabled, or nil otherwise.
func (c *Consumer) Logs() chan LogEvent {
	return nil
}

// ReadMessage polls the consumer for a message.
//
// This is a convenience API that wraps Poll() and only returns
// messages or errors. All other event types are discarded.
//
// The call will block for at most `timeout` waiting for
// a new message or error. `timeout` may be set to -1 for
// indefinite wait.
//
// Timeout is returned as (nil, err) where err is `kafka.(Error).Code == Kafka.ErrTimedOut`.
//
// Messages are returned as (msg, nil),
// while general errors are returned as (nil, err),
// and partition-specific errors are returned as (msg, err) where
// msg.TopicPartition provides partition-specific information (such as topic, partition and offset).
//
// All other event types, such as PartitionEOF, AssignedPartitions, etc, are silently discarded.
func (c *Consumer) ReadMessage(timeout time.Duration) (*Message, error) {
	return nil, errors.New("")
}

// Close Consumer instance.
// The object is no longer usable after this call.
func (c *Consumer) Close() (err error) {
	return errors.New("")
}

// NewConsumer creates a new high-level Consumer instance.
//
// conf is a *ConfigMap with standard librdkafka configuration properties.
//
// Supported special configuration properties:
//
//	go.application.rebalance.enable (bool, false) - Forward rebalancing responsibility to application via the Events() channel.
//	                                     If set to true the app must handle the AssignedPartitions and
//	                                     RevokedPartitions events and call Assign() and Unassign()
//	                                     respectively.
//	go.events.channel.enable (bool, false) - Enable the Events() channel. Messages and events will be pushed on the Events() channel and the Poll() interface will be disabled. (Experimental)
//	go.events.channel.size (int, 1000) - Events() channel size
//	go.logs.channel.enable (bool, false) - Forward log to Logs() channel.
//	go.logs.channel (chan kafka.LogEvent, nil) - Forward logs to application-provided channel instead of Logs(). Requires go.logs.channel.enable=true.
//
// WARNING: Due to the buffering nature of channels (and queues in general) the
// use of the events channel risks receiving outdated events and
// messages. Minimizing go.events.channel.size reduces the risk
// and number of outdated events and messages but does not eliminate
// the factor completely. With a channel size of 1 at most one
// event or message may be outdated.
func NewConsumer(conf *ConfigMap) (*Consumer, error) {
	return nil, errors.New("")
}

// GetMetadata queries broker for cluster and topic metadata.
// If topic is non-nil only information about that topic is returned, else if
// allTopics is false only information about locally used topics is returned,
// else information about all topics is returned.
// GetMetadata is equivalent to listTopics, describeTopics and describeCluster in the Java API.
func (c *Consumer) GetMetadata(topic *string, allTopics bool, timeoutMs int) (*Metadata, error) {
	return nil, errors.New("")
}

// QueryWatermarkOffsets queries the broker for the low and high offsets for the given topic and partition.
func (c *Consumer) QueryWatermarkOffsets(topic string, partition int32, timeoutMs int) (low, high int64, err error) {
	return 0, 0, errors.New("")
}

// GetWatermarkOffsets returns the cached low and high offsets for the given topic
// and partition.  The high offset is populated on every fetch response or via calling QueryWatermarkOffsets.
// The low offset is populated every statistics.interval.ms if that value is set.
// OffsetInvalid will be returned if there is no cached offset for either value.
func (c *Consumer) GetWatermarkOffsets(topic string, partition int32) (low, high int64, err error) {
	return 0, 0, errors.New("")
}

// OffsetsForTimes looks up offsets by timestamp for the given partitions.
//
// The returned offset for each partition is the earliest offset whose
// timestamp is greater than or equal to the given timestamp in the
// corresponding partition. If the provided timestamp exceeds that of the
// last message in the partition, a value of -1 will be returned.
//
// The timestamps to query are represented as `.Offset` in the `times`
// argument and the looked up offsets are represented as `.Offset` in the returned
// `offsets` list.
//
// The function will block for at most timeoutMs milliseconds.
//
// Duplicate Topic+Partitions are not supported.
// Per-partition errors may be returned in the `.Error` field.
func (c *Consumer) OffsetsForTimes(times []TopicPartition, timeoutMs int) (offsets []TopicPartition, err error) {
	return nil, errors.New("")
}

// Subscription returns the current subscription as set by Subscribe()
func (c *Consumer) Subscription() (topics []string, err error) {
	return nil, errors.New("")
}

// Assignment returns the current partition assignments
func (c *Consumer) Assignment() (partitions []TopicPartition, err error) {
	return nil, errors.New("")
}

// Committed retrieves committed offsets for the given set of partitions
func (c *Consumer) Committed(partitions []TopicPartition, timeoutMs int) (offsets []TopicPartition, err error) {
	return nil, errors.New("")
}

// Position returns the current consume position for the given partitions.
// Typical use is to call Assignment() to get the partition list
// and then pass it to Position() to get the current consume position for
// each of the assigned partitions.
// The conusme position is the next message to read from the partition.
// i.e., the offset of the last message seen by the application + 1.
func (c *Consumer) Position(partitions []TopicPartition) (offsets []TopicPartition, err error) {
	return nil, errors.New("")
}

// Pause consumption for the provided list of partitions
//
// Note that messages already enqueued on the consumer's Event channel
// (if `go.events.channel.enable` has been set) will NOT be purged by
// this call, set `go.events.channel.size` accordingly.
func (c *Consumer) Pause(partitions []TopicPartition) (err error) {
	return errors.New("")
}

// Resume consumption for the provided list of partitions
func (c *Consumer) Resume(partitions []TopicPartition) (err error) {
	return errors.New("")
}

// SetOAuthBearerToken sets the the data to be transmitted
// to a broker during SASL/OAUTHBEARER authentication. It will return nil
// on success, otherwise an error if:
// 1) the token data is invalid (meaning an expiration time in the past
// or either a token value or an extension key or value that does not meet
// the regular expression requirements as per
// https://tools.ietf.org/html/rfc7628#section-3.1);
// 2) SASL/OAUTHBEARER is not supported by the underlying librdkafka build;
// 3) SASL/OAUTHBEARER is supported but is not configured as the client's
// authentication mechanism.
func (c *Consumer) SetOAuthBearerToken(oauthBearerToken OAuthBearerToken) error {
	return errors.New("")
}

// SetOAuthBearerTokenFailure sets the error message describing why token
// retrieval/setting failed; it also schedules a new token refresh event for 10
// seconds later so the attempt may be retried. It will return nil on
// success, otherwise an error if:
// 1) SASL/OAUTHBEARER is not supported by the underlying librdkafka build;
// 2) SASL/OAUTHBEARER is supported but is not configured as the client's
// authentication mechanism.
func (c *Consumer) SetOAuthBearerTokenFailure(errstr string) error {
	return errors.New("")
}

// ConsumerGroupMetadata reflects the current consumer group member metadata.
type ConsumerGroupMetadata struct {
	serialized []byte
}

// GetConsumerGroupMetadata returns the consumer's current group metadata.
// This object should be passed to the transactional producer's
// SendOffsetsToTransaction() API.
func (c *Consumer) GetConsumerGroupMetadata() (*ConsumerGroupMetadata, error) {
	return nil, errors.New("")
}

// NewTestConsumerGroupMetadata creates a new consumer group metadata instance
// mainly for testing use.
// Use GetConsumerGroupMetadata() to retrieve the real metadata.
func NewTestConsumerGroupMetadata(groupID string) (*ConsumerGroupMetadata, error) {
	return nil, errors.New("")
}
