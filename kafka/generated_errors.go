package kafka

// Copyright 2016-2020 Confluent Inc.
// AUTOMATICALLY GENERATED ON 2020-04-08 13:17:23.310983 +0200 CEST m=+0.001902312 USING librdkafka 1.4.0

// ErrorCode is the integer representation of local and broker error codes
type ErrorCode int

// String returns a human readable representation of an error code
func (c ErrorCode) String() string {
	return ""
}

const (
	// ErrBadMsg Local: Bad message format
	ErrBadMsg ErrorCode = ErrorCode(0)
	// ErrBadCompression Local: Invalid compressed data
	ErrBadCompression ErrorCode = ErrorCode(0)
	// ErrDestroy Local: Broker handle destroyed
	ErrDestroy ErrorCode = ErrorCode(0)
	// ErrFail Local: Communication failure with broker
	ErrFail ErrorCode = ErrorCode(0)
	// ErrTransport Local: Broker transport failure
	ErrTransport ErrorCode = ErrorCode(0)
	// ErrCritSysResource Local: Critical system resource failure
	ErrCritSysResource ErrorCode = ErrorCode(0)
	// ErrResolve Local: Host resolution failure
	ErrResolve ErrorCode = ErrorCode(0)
	// ErrMsgTimedOut Local: Message timed out
	ErrMsgTimedOut ErrorCode = ErrorCode(0)
	// ErrPartitionEOF Broker: No more messages
	ErrPartitionEOF ErrorCode = ErrorCode(0)
	// ErrUnknownPartition Local: Unknown partition
	ErrUnknownPartition ErrorCode = ErrorCode(0)
	// ErrFs Local: File or filesystem error
	ErrFs ErrorCode = ErrorCode(0)
	// ErrUnknownTopic Local: Unknown topic
	ErrUnknownTopic ErrorCode = ErrorCode(0)
	// ErrAllBrokersDown Local: All broker connections are down
	ErrAllBrokersDown ErrorCode = ErrorCode(0)
	// ErrInvalidArg Local: Invalid argument or configuration
	ErrInvalidArg ErrorCode = ErrorCode(0)
	// ErrTimedOut Local: Timed out
	ErrTimedOut ErrorCode = ErrorCode(0)
	// ErrQueueFull Local: Queue full
	ErrQueueFull ErrorCode = ErrorCode(0)
	// ErrIsrInsuff Local: ISR count insufficient
	ErrIsrInsuff ErrorCode = ErrorCode(0)
	// ErrNodeUpdate Local: Broker node update
	ErrNodeUpdate ErrorCode = ErrorCode(0)
	// ErrSsl Local: SSL error
	ErrSsl ErrorCode = ErrorCode(0)
	// ErrWaitCoord Local: Waiting for coordinator
	ErrWaitCoord ErrorCode = ErrorCode(0)
	// ErrUnknownGroup Local: Unknown group
	ErrUnknownGroup ErrorCode = ErrorCode(0)
	// ErrInProgress Local: Operation in progress
	ErrInProgress ErrorCode = ErrorCode(0)
	// ErrPrevInProgress Local: Previous operation in progress
	ErrPrevInProgress ErrorCode = ErrorCode(0)
	// ErrExistingSubscription Local: Existing subscription
	ErrExistingSubscription ErrorCode = ErrorCode(0)
	// ErrAssignPartitions Local: Assign partitions
	ErrAssignPartitions ErrorCode = ErrorCode(0)
	// ErrRevokePartitions Local: Revoke partitions
	ErrRevokePartitions ErrorCode = ErrorCode(0)
	// ErrConflict Local: Conflicting use
	ErrConflict ErrorCode = ErrorCode(0)
	// ErrState Local: Erroneous state
	ErrState ErrorCode = ErrorCode(0)
	// ErrUnknownProtocol Local: Unknown protocol
	ErrUnknownProtocol ErrorCode = ErrorCode(0)
	// ErrNotImplemented Local: Not implemented
	ErrNotImplemented ErrorCode = ErrorCode(0)
	// ErrAuthentication Local: Authentication failure
	ErrAuthentication ErrorCode = ErrorCode(0)
	// ErrNoOffset Local: No offset stored
	ErrNoOffset ErrorCode = ErrorCode(0)
	// ErrOutdated Local: Outdated
	ErrOutdated ErrorCode = ErrorCode(0)
	// ErrTimedOutQueue Local: Timed out in queue
	ErrTimedOutQueue ErrorCode = ErrorCode(0)
	// ErrUnsupportedFeature Local: Required feature not supported by broker
	ErrUnsupportedFeature ErrorCode = ErrorCode(0)
	// ErrWaitCache Local: Awaiting cache update
	ErrWaitCache ErrorCode = ErrorCode(0)
	// ErrIntr Local: Operation interrupted
	ErrIntr ErrorCode = ErrorCode(0)
	// ErrKeySerialization Local: Key serialization error
	ErrKeySerialization ErrorCode = ErrorCode(0)
	// ErrValueSerialization Local: Value serialization error
	ErrValueSerialization ErrorCode = ErrorCode(0)
	// ErrKeyDeserialization Local: Key deserialization error
	ErrKeyDeserialization ErrorCode = ErrorCode(0)
	// ErrValueDeserialization Local: Value deserialization error
	ErrValueDeserialization ErrorCode = ErrorCode(0)
	// ErrPartial Local: Partial response
	ErrPartial ErrorCode = ErrorCode(0)
	// ErrReadOnly Local: Read-only object
	ErrReadOnly ErrorCode = ErrorCode(0)
	// ErrNoent Local: No such entry
	ErrNoent ErrorCode = ErrorCode(0)
	// ErrUnderflow Local: Read underflow
	ErrUnderflow ErrorCode = ErrorCode(0)
	// ErrInvalidType Local: Invalid type
	ErrInvalidType ErrorCode = ErrorCode(0)
	// ErrRetry Local: Retry operation
	ErrRetry ErrorCode = ErrorCode(0)
	// ErrPurgeQueue Local: Purged in queue
	ErrPurgeQueue ErrorCode = ErrorCode(0)
	// ErrPurgeInflight Local: Purged in flight
	ErrPurgeInflight ErrorCode = ErrorCode(0)
	// ErrFatal Local: Fatal error
	ErrFatal ErrorCode = ErrorCode(0)
	// ErrInconsistent Local: Inconsistent state
	ErrInconsistent ErrorCode = ErrorCode(0)
	// ErrGaplessGuarantee Local: Gap-less ordering would not be guaranteed if proceeding
	ErrGaplessGuarantee ErrorCode = ErrorCode(0)
	// ErrMaxPollExceeded Local: Maximum application poll interval (max.poll.interval.ms) exceeded
	ErrMaxPollExceeded ErrorCode = ErrorCode(0)
	// ErrUnknownBroker Local: Unknown broker
	ErrUnknownBroker ErrorCode = ErrorCode(0)
	// ErrNotConfigured Local: Functionality not configured
	ErrNotConfigured ErrorCode = ErrorCode(0)
	// ErrFenced Local: This instance has been fenced by a newer instance
	ErrFenced ErrorCode = ErrorCode(0)
	// ErrApplication Local: Application generated error
	ErrApplication ErrorCode = ErrorCode(0)
	// ErrUnknown Unknown broker error
	ErrUnknown ErrorCode = ErrorCode(0)
	// ErrNoError Success
	ErrNoError ErrorCode = ErrorCode(0)
	// ErrOffsetOutOfRange Broker: Offset out of range
	ErrOffsetOutOfRange ErrorCode = ErrorCode(0)
	// ErrInvalidMsg Broker: Invalid message
	ErrInvalidMsg ErrorCode = ErrorCode(0)
	// ErrUnknownTopicOrPart Broker: Unknown topic or partition
	ErrUnknownTopicOrPart ErrorCode = ErrorCode(0)
	// ErrInvalidMsgSize Broker: Invalid message size
	ErrInvalidMsgSize ErrorCode = ErrorCode(0)
	// ErrLeaderNotAvailable Broker: Leader not available
	ErrLeaderNotAvailable ErrorCode = ErrorCode(0)
	// ErrNotLeaderForPartition Broker: Not leader for partition
	ErrNotLeaderForPartition ErrorCode = ErrorCode(0)
	// ErrRequestTimedOut Broker: Request timed out
	ErrRequestTimedOut ErrorCode = ErrorCode(0)
	// ErrBrokerNotAvailable Broker: Broker not available
	ErrBrokerNotAvailable ErrorCode = ErrorCode(0)
	// ErrReplicaNotAvailable Broker: Replica not available
	ErrReplicaNotAvailable ErrorCode = ErrorCode(0)
	// ErrMsgSizeTooLarge Broker: Message size too large
	ErrMsgSizeTooLarge ErrorCode = ErrorCode(0)
	// ErrStaleCtrlEpoch Broker: StaleControllerEpochCode
	ErrStaleCtrlEpoch ErrorCode = ErrorCode(0)
	// ErrOffsetMetadataTooLarge Broker: Offset metadata string too large
	ErrOffsetMetadataTooLarge ErrorCode = ErrorCode(0)
	// ErrNetworkException Broker: Broker disconnected before response received
	ErrNetworkException ErrorCode = ErrorCode(0)
	// ErrCoordinatorLoadInProgress Broker: Coordinator load in progress
	ErrCoordinatorLoadInProgress ErrorCode = ErrorCode(0)
	// ErrCoordinatorNotAvailable Broker: Coordinator not available
	ErrCoordinatorNotAvailable ErrorCode = ErrorCode(0)
	// ErrNotCoordinator Broker: Not coordinator
	ErrNotCoordinator ErrorCode = ErrorCode(0)
	// ErrTopicException Broker: Invalid topic
	ErrTopicException ErrorCode = ErrorCode(0)
	// ErrRecordListTooLarge Broker: Message batch larger than configured server segment size
	ErrRecordListTooLarge ErrorCode = ErrorCode(0)
	// ErrNotEnoughReplicas Broker: Not enough in-sync replicas
	ErrNotEnoughReplicas ErrorCode = ErrorCode(0)
	// ErrNotEnoughReplicasAfterAppend Broker: Message(s) written to insufficient number of in-sync replicas
	ErrNotEnoughReplicasAfterAppend ErrorCode = ErrorCode(0)
	// ErrInvalidRequiredAcks Broker: Invalid required acks value
	ErrInvalidRequiredAcks ErrorCode = ErrorCode(0)
	// ErrIllegalGeneration Broker: Specified group generation id is not valid
	ErrIllegalGeneration ErrorCode = ErrorCode(0)
	// ErrInconsistentGroupProtocol Broker: Inconsistent group protocol
	ErrInconsistentGroupProtocol ErrorCode = ErrorCode(0)
	// ErrInvalidGroupID Broker: Invalid group.id
	ErrInvalidGroupID ErrorCode = ErrorCode(0)
	// ErrUnknownMemberID Broker: Unknown member
	ErrUnknownMemberID ErrorCode = ErrorCode(0)
	// ErrInvalidSessionTimeout Broker: Invalid session timeout
	ErrInvalidSessionTimeout ErrorCode = ErrorCode(0)
	// ErrRebalanceInProgress Broker: Group rebalance in progress
	ErrRebalanceInProgress ErrorCode = ErrorCode(0)
	// ErrInvalidCommitOffsetSize Broker: Commit offset data size is not valid
	ErrInvalidCommitOffsetSize ErrorCode = ErrorCode(0)
	// ErrTopicAuthorizationFailed Broker: Topic authorization failed
	ErrTopicAuthorizationFailed ErrorCode = ErrorCode(0)
	// ErrGroupAuthorizationFailed Broker: Group authorization failed
	ErrGroupAuthorizationFailed ErrorCode = ErrorCode(0)
	// ErrClusterAuthorizationFailed Broker: Cluster authorization failed
	ErrClusterAuthorizationFailed ErrorCode = ErrorCode(0)
	// ErrInvalidTimestamp Broker: Invalid timestamp
	ErrInvalidTimestamp ErrorCode = ErrorCode(0)
	// ErrUnsupportedSaslMechanism Broker: Unsupported SASL mechanism
	ErrUnsupportedSaslMechanism ErrorCode = ErrorCode(0)
	// ErrIllegalSaslState Broker: Request not valid in current SASL state
	ErrIllegalSaslState ErrorCode = ErrorCode(0)
	// ErrUnsupportedVersion Broker: API version not supported
	ErrUnsupportedVersion ErrorCode = ErrorCode(0)
	// ErrTopicAlreadyExists Broker: Topic already exists
	ErrTopicAlreadyExists ErrorCode = ErrorCode(0)
	// ErrInvalidPartitions Broker: Invalid number of partitions
	ErrInvalidPartitions ErrorCode = ErrorCode(0)
	// ErrInvalidReplicationFactor Broker: Invalid replication factor
	ErrInvalidReplicationFactor ErrorCode = ErrorCode(0)
	// ErrInvalidReplicaAssignment Broker: Invalid replica assignment
	ErrInvalidReplicaAssignment ErrorCode = ErrorCode(0)
	// ErrInvalidConfig Broker: Configuration is invalid
	ErrInvalidConfig ErrorCode = ErrorCode(0)
	// ErrNotController Broker: Not controller for cluster
	ErrNotController ErrorCode = ErrorCode(0)
	// ErrInvalidRequest Broker: Invalid request
	ErrInvalidRequest ErrorCode = ErrorCode(0)
	// ErrUnsupportedForMessageFormat Broker: Message format on broker does not support request
	ErrUnsupportedForMessageFormat ErrorCode = ErrorCode(0)
	// ErrPolicyViolation Broker: Policy violation
	ErrPolicyViolation ErrorCode = ErrorCode(0)
	// ErrOutOfOrderSequenceNumber Broker: Broker received an out of order sequence number
	ErrOutOfOrderSequenceNumber ErrorCode = ErrorCode(0)
	// ErrDuplicateSequenceNumber Broker: Broker received a duplicate sequence number
	ErrDuplicateSequenceNumber ErrorCode = ErrorCode(0)
	// ErrInvalidProducerEpoch Broker: Producer attempted an operation with an old epoch
	ErrInvalidProducerEpoch ErrorCode = ErrorCode(0)
	// ErrInvalidTxnState Broker: Producer attempted a transactional operation in an invalid state
	ErrInvalidTxnState ErrorCode = ErrorCode(0)
	// ErrInvalidProducerIDMapping Broker: Producer attempted to use a producer id which is not currently assigned to its transactional id
	ErrInvalidProducerIDMapping ErrorCode = ErrorCode(0)
	// ErrInvalidTransactionTimeout Broker: Transaction timeout is larger than the maximum value allowed by the broker's max.transaction.timeout.ms
	ErrInvalidTransactionTimeout ErrorCode = ErrorCode(0)
	// ErrConcurrentTransactions Broker: Producer attempted to update a transaction while another concurrent operation on the same transaction was ongoing
	ErrConcurrentTransactions ErrorCode = ErrorCode(0)
	// ErrTransactionCoordinatorFenced Broker: Indicates that the transaction coordinator sending a WriteTxnMarker is no longer the current coordinator for a given producer
	ErrTransactionCoordinatorFenced ErrorCode = ErrorCode(0)
	// ErrTransactionalIDAuthorizationFailed Broker: Transactional Id authorization failed
	ErrTransactionalIDAuthorizationFailed ErrorCode = ErrorCode(0)
	// ErrSecurityDisabled Broker: Security features are disabled
	ErrSecurityDisabled ErrorCode = ErrorCode(0)
	// ErrOperationNotAttempted Broker: Operation not attempted
	ErrOperationNotAttempted ErrorCode = ErrorCode(0)
	// ErrKafkaStorageError Broker: Disk error when trying to access log file on disk
	ErrKafkaStorageError ErrorCode = ErrorCode(0)
	// ErrLogDirNotFound Broker: The user-specified log directory is not found in the broker config
	ErrLogDirNotFound ErrorCode = ErrorCode(0)
	// ErrSaslAuthenticationFailed Broker: SASL Authentication failed
	ErrSaslAuthenticationFailed ErrorCode = ErrorCode(0)
	// ErrUnknownProducerID Broker: Unknown Producer Id
	ErrUnknownProducerID ErrorCode = ErrorCode(0)
	// ErrReassignmentInProgress Broker: Partition reassignment is in progress
	ErrReassignmentInProgress ErrorCode = ErrorCode(0)
	// ErrDelegationTokenAuthDisabled Broker: Delegation Token feature is not enabled
	ErrDelegationTokenAuthDisabled ErrorCode = ErrorCode(0)
	// ErrDelegationTokenNotFound Broker: Delegation Token is not found on server
	ErrDelegationTokenNotFound ErrorCode = ErrorCode(0)
	// ErrDelegationTokenOwnerMismatch Broker: Specified Principal is not valid Owner/Renewer
	ErrDelegationTokenOwnerMismatch ErrorCode = ErrorCode(0)
	// ErrDelegationTokenRequestNotAllowed Broker: Delegation Token requests are not allowed on this connection
	ErrDelegationTokenRequestNotAllowed ErrorCode = ErrorCode(0)
	// ErrDelegationTokenAuthorizationFailed Broker: Delegation Token authorization failed
	ErrDelegationTokenAuthorizationFailed ErrorCode = ErrorCode(0)
	// ErrDelegationTokenExpired Broker: Delegation Token is expired
	ErrDelegationTokenExpired ErrorCode = ErrorCode(0)
	// ErrInvalidPrincipalType Broker: Supplied principalType is not supported
	ErrInvalidPrincipalType ErrorCode = ErrorCode(0)
	// ErrNonEmptyGroup Broker: The group is not empty
	ErrNonEmptyGroup ErrorCode = ErrorCode(0)
	// ErrGroupIDNotFound Broker: The group id does not exist
	ErrGroupIDNotFound ErrorCode = ErrorCode(0)
	// ErrFetchSessionIDNotFound Broker: The fetch session ID was not found
	ErrFetchSessionIDNotFound ErrorCode = ErrorCode(0)
	// ErrInvalidFetchSessionEpoch Broker: The fetch session epoch is invalid
	ErrInvalidFetchSessionEpoch ErrorCode = ErrorCode(0)
	// ErrListenerNotFound Broker: No matching listener
	ErrListenerNotFound ErrorCode = ErrorCode(0)
	// ErrTopicDeletionDisabled Broker: Topic deletion is disabled
	ErrTopicDeletionDisabled ErrorCode = ErrorCode(0)
	// ErrFencedLeaderEpoch Broker: Leader epoch is older than broker epoch
	ErrFencedLeaderEpoch ErrorCode = ErrorCode(0)
	// ErrUnknownLeaderEpoch Broker: Leader epoch is newer than broker epoch
	ErrUnknownLeaderEpoch ErrorCode = ErrorCode(0)
	// ErrUnsupportedCompressionType Broker: Unsupported compression type
	ErrUnsupportedCompressionType ErrorCode = ErrorCode(0)
	// ErrStaleBrokerEpoch Broker: Broker epoch has changed
	ErrStaleBrokerEpoch ErrorCode = ErrorCode(0)
	// ErrOffsetNotAvailable Broker: Leader high watermark is not caught up
	ErrOffsetNotAvailable ErrorCode = ErrorCode(0)
	// ErrMemberIDRequired Broker: Group member needs a valid member ID
	ErrMemberIDRequired ErrorCode = ErrorCode(0)
	// ErrPreferredLeaderNotAvailable Broker: Preferred leader was not available
	ErrPreferredLeaderNotAvailable ErrorCode = ErrorCode(0)
	// ErrGroupMaxSizeReached Broker: Consumer group has reached maximum size
	ErrGroupMaxSizeReached ErrorCode = ErrorCode(0)
	// ErrFencedInstanceID Broker: Static consumer fenced by other consumer with same group.instance.id
	ErrFencedInstanceID ErrorCode = ErrorCode(0)
)
