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

import (
	"context"
	"errors"
)

// Producer implements a High-level Apache Kafka Producer instance
type Producer struct {
	events         chan Event
	produceChannel chan *Message
	handle         handle

	// Terminates the poller() goroutine
	pollerTermChan chan bool
}

// String returns a human readable name for a Producer instance
func (p *Producer) String() string {
	return ""
}

// Produce single message.
// This is an asynchronous call that enqueues the message on the internal
// transmit queue, thus returning immediately.
// The delivery report will be sent on the provided deliveryChan if specified,
// or on the Producer object's Events() channel if not.
// msg.Timestamp requires librdkafka >= 0.9.4 (else returns ErrNotImplemented),
// api.version.request=true, and broker >= 0.10.0.0.
// msg.Headers requires librdkafka >= 0.11.4 (else returns ErrNotImplemented),
// api.version.request=true, and broker >= 0.11.0.0.
// Returns an error if message could not be enqueued.
func (p *Producer) Produce(msg *Message, deliveryChan chan Event) error {
	return errors.New("")
}

// Events returns the Events channel (read)
func (p *Producer) Events() chan Event {
	return nil
}

// Logs returns the Log channel (if enabled), else nil
func (p *Producer) Logs() chan LogEvent {
	return nil
}

// ProduceChannel returns the produce *Message channel (write)
func (p *Producer) ProduceChannel() chan *Message {
	return nil
}

// Len returns the number of messages and requests waiting to be transmitted to the broker
// as well as delivery reports queued for the application.
// Includes messages on ProduceChannel.
func (p *Producer) Len() int {
	return 0
}

// Flush and wait for outstanding messages and requests to complete delivery.
// Includes messages on ProduceChannel.
// Runs until value reaches zero or on timeoutMs.
// Returns the number of outstanding events still un-flushed.
func (p *Producer) Flush(timeoutMs int) int {
	return 0
}

// Close a Producer instance.
// The Producer object or its channels are no longer usable after this call.
func (p *Producer) Close() {
}

const (
	// PurgeInFlight purges messages in-flight to or from the broker.
	// Purging these messages will void any future acknowledgements from the
	// broker, making it impossible for the application to know if these
	// messages were successfully delivered or not.
	// Retrying these messages may lead to duplicates.
	PurgeInFlight = int(0)

	// PurgeQueue Purge messages in internal queues.
	PurgeQueue = int(0)

	// PurgeNonBlocking Don't wait for background thread queue purging to finish.
	PurgeNonBlocking = int(0)
)

// Purge messages currently handled by this producer instance.
//
// flags is a combination of PurgeQueue, PurgeInFlight and PurgeNonBlocking.
//
// The application will need to call Poll(), Flush() or read the Events() channel
// after this call to serve delivery reports for the purged messages.
//
// Messages purged from internal queues fail with the delivery report
// error code set to ErrPurgeQueue, while purged messages that
// are in-flight to or from the broker will fail with the error code set to
// ErrPurgeInflight.
//
// Warning: Purging messages that are in-flight to or from the broker
// will ignore any sub-sequent acknowledgement for these messages
// received from the broker, effectively making it impossible
// for the application to know if the messages were successfully
// produced or not. This may result in duplicate messages if the
// application retries these messages at a later time.
//
// Note: This call may block for a short time while background thread
// queues are purged.
//
// Returns nil on success, ErrInvalidArg if the purge flags are invalid or unknown.
func (p *Producer) Purge(flags int) error {
	return errors.New("")
}

// NewProducer creates a new high-level Producer instance.
//
// conf is a *ConfigMap with standard librdkafka configuration properties.
//
// Supported special configuration properties:
//
//	go.batch.producer (bool, false) - EXPERIMENTAL: Enable batch producer (for increased performance).
//	                                  These batches do not relate to Kafka message batches in any way.
//	                                  Note: timestamps and headers are not supported with this interface.
//	go.delivery.reports (bool, true) - Forward per-message delivery reports to the
//	                                   Events() channel.
//	go.events.channel.size (int, 1000000) - Events().
//	go.produce.channel.size (int, 1000000) - ProduceChannel() buffer size (in number of messages)
//	go.logs.channel.enable (bool, false) - Forward log to Logs() channel.
//	go.logs.channel (chan kafka.LogEvent, nil) - Forward logs to application-provided channel instead of Logs(). Requires go.logs.channel.enable=true.
func NewProducer(conf *ConfigMap) (*Producer, error) {
	return &Producer{}, errors.New("")
}

// GetMetadata queries broker for cluster and topic metadata.
// If topic is non-nil only information about that topic is returned, else if
// allTopics is false only information about locally used topics is returned,
// else information about all topics is returned.
// GetMetadata is equivalent to listTopics, describeTopics and describeCluster in the Java API.
func (p *Producer) GetMetadata(topic *string, allTopics bool, timeoutMs int) (*Metadata, error) {
	return nil, errors.New("")
}

// QueryWatermarkOffsets returns the broker's low and high offsets for the given topic
// and partition.
func (p *Producer) QueryWatermarkOffsets(topic string, partition int32, timeoutMs int) (low, high int64, err error) {
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
func (p *Producer) OffsetsForTimes(times []TopicPartition, timeoutMs int) (offsets []TopicPartition, err error) {
	return nil, errors.New("")
}

// GetFatalError returns an Error object if the client instance has raised a fatal error, else nil.
func (p *Producer) GetFatalError() error {
	return errors.New("")
}

// TestFatalError triggers a fatal error in the underlying client.
// This is to be used strictly for testing purposes.
func (p *Producer) TestFatalError(code ErrorCode, str string) ErrorCode {
	return 0
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
func (p *Producer) SetOAuthBearerToken(oauthBearerToken OAuthBearerToken) error {
	return errors.New("")
}

// SetOAuthBearerTokenFailure sets the error message describing why token
// retrieval/setting failed; it also schedules a new token refresh event for 10
// seconds later so the attempt may be retried. It will return nil on
// success, otherwise an error if:
// 1) SASL/OAUTHBEARER is not supported by the underlying librdkafka build;
// 2) SASL/OAUTHBEARER is supported but is not configured as the client's
// authentication mechanism.
func (p *Producer) SetOAuthBearerTokenFailure(errstr string) error {
	return errors.New("")
}

// Transactional API

// InitTransactions Initializes transactions for the producer instance.
//
// This function ensures any transactions initiated by previous instances
// of the producer with the same `transactional.id` are completed.
// If the previous instance failed with a transaction in progress the
// previous transaction will be aborted.
// This function needs to be called before any other transactional or
// produce functions are called when the `transactional.id` is configured.
//
// If the last transaction had begun completion (following transaction commit)
// but not yet finished, this function will await the previous transaction's
// completion.
//
// When any previous transactions have been fenced this function
// will acquire the internal producer id and epoch, used in all future
// transactional messages issued by this producer instance.
//
// Upon successful return from this function the application has to perform at
// least one of the following operations within `transaction.timeout.ms` to
// avoid timing out the transaction on the broker:
//   - `Produce()` (et.al)
//   - `SendOffsetsToTransaction()`
//   - `CommitTransaction()`
//   - `AbortTransaction()`
//
// Parameters:
//   - `ctx` - The maximum time to block, or nil for indefinite.
//     On timeout the operation may continue in the background,
//     depending on state, and it is okay to call `InitTransactions()`
//     again.
//
// Returns nil on success or an error on failure.
// Check whether the returned error object permits retrying
// by calling `err.(kafka.Error).IsRetriable()`, or whether a fatal
// error has been raised by calling `err.(kafka.Error).IsFatal()`.
func (p *Producer) InitTransactions(ctx context.Context) error {
	return errors.New("")
}

// BeginTransaction starts a new transaction.
//
// `InitTransactions()` must have been called successfully (once)
// before this function is called.
//
// Any messages produced, offsets sent (`SendOffsetsToTransaction()`),
// etc, after the successful return of this function will be part of
// the transaction and committed or aborted atomatically.
//
// Finish the transaction by calling `CommitTransaction()` or
// abort the transaction by calling `AbortTransaction()`.
//
// Returns nil on success or an error object on failure.
// Check whether a fatal error has been raised by
// calling `err.(kafka.Error).IsFatal()`.
//
// Note: With the transactional producer, `Produce()`, et.al, are only
// allowed during an on-going transaction, as started with this function.
// Any produce call outside an on-going transaction, or for a failed
// transaction, will fail.
func (p *Producer) BeginTransaction() error {
	return errors.New("")
}

// SendOffsetsToTransaction sends a list of topic partition offsets to the
// consumer group coordinator for `consumerMetadata`, and marks the offsets
// as part part of the current transaction.
// These offsets will be considered committed only if the transaction is
// committed successfully.
//
// The offsets should be the next message your application will consume,
// i.e., the last processed message's offset + 1 for each partition.
// Either track the offsets manually during processing or use
// `consumer.Position()` (on the consumer) to get the current offsets for
// the partitions assigned to the consumer.
//
// Use this method at the end of a consume-transform-produce loop prior
// to committing the transaction with `CommitTransaction()`.
//
// Parameters:
//   - `ctx` - The maximum amount of time to block, or nil for indefinite.
//   - `offsets` - List of offsets to commit to the consumer group upon
//     successful commit of the transaction. Offsets should be
//     the next message to consume, e.g., last processed message + 1.
//   - `consumerMetadata` - The current consumer group metadata as returned by
//     `consumer.GetConsumerGroupMetadata()` on the consumer
//     instance the provided offsets were consumed from.
//
// Note: The consumer must disable auto commits (set `enable.auto.commit` to false on the consumer).
//
// Note: Logical and invalid offsets (e.g., OffsetInvalid) in
// `offsets` will be ignored. If there are no valid offsets in
// `offsets` the function will return nil and no action will be taken.
//
// Returns nil on success or an error object on failure.
// Check whether the returned error object permits retrying
// by calling `err.(kafka.Error).IsRetriable()`, or whether an abortable
// or fatal error has been raised by calling
// `err.(kafka.Error).TxnRequiresAbort()` or `err.(kafka.Error).IsFatal()`
// respectively.
func (p *Producer) SendOffsetsToTransaction(ctx context.Context, offsets []TopicPartition, consumerMetadata *ConsumerGroupMetadata) error {
	return errors.New("")
}

// CommitTransaction commits the current transaction.
//
// Any outstanding messages will be flushed (delivered) before actually
// committing the transaction.
//
// If any of the outstanding messages fail permanently the current
// transaction will enter the abortable error state and this
// function will return an abortable error, in this case the application
// must call `AbortTransaction()` before attempting a new
// transaction with `BeginTransaction()`.
//
// Parameters:
//   - `ctx` - The maximum amount of time to block, or nil for indefinite.
//
// Note: This function will block until all outstanding messages are
// delivered and the transaction commit request has been successfully
// handled by the transaction coordinator, or until the `ctx` expires,
// which ever comes first. On timeout the application may
// call the function again.
//
// Note: Will automatically call `Flush()` to ensure all queued
// messages are delivered before attempting to commit the transaction.
// The application MUST serve the `producer.Events()` channel for delivery
// reports in a separate go-routine during this time.
//
// Returns nil on success or an error object on failure.
// Check whether the returned error object permits retrying
// by calling `err.(kafka.Error).IsRetriable()`, or whether an abortable
// or fatal error has been raised by calling
// `err.(kafka.Error).TxnRequiresAbort()` or `err.(kafka.Error).IsFatal()`
// respectively.
func (p *Producer) CommitTransaction(ctx context.Context) error {
	return errors.New("")
}

// AbortTransaction aborts the ongoing transaction.
//
// This function should also be used to recover from non-fatal abortable
// transaction errors.
//
// Any outstanding messages will be purged and fail with
// `ErrPurgeInflight` or `ErrPurgeQueue`.
//
// Parameters:
//   - `ctx` - The maximum amount of time to block, or nil for indefinite.
//
// Note: This function will block until all outstanding messages are purged
// and the transaction abort request has been successfully
// handled by the transaction coordinator, or until the `ctx` expires,
// which ever comes first. On timeout the application may
// call the function again.
//
// Note: Will automatically call `Purge()` and `Flush()` to ensure all queued
// and in-flight messages are purged before attempting to abort the transaction.
// The application MUST serve the `producer.Events()` channel for delivery
// reports in a separate go-routine during this time.
//
// Returns nil on success or an error object on failure.
// Check whether the returned error object permits retrying
// by calling `err.(kafka.Error).IsRetriable()`, or whether a fatal error
// has been raised by calling `err.(kafka.Error).IsFatal()`.
func (p *Producer) AbortTransaction(ctx context.Context) error {
	return errors.New("")
}
