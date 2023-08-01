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

package kafka

import (
	"context"
	"errors"
)

// AdminClient is derived from an existing Producer or Consumer
type AdminClient struct {
	handle    *handle
	isDerived bool // Derived from existing client handle
}

// TopicResult provides per-topic operation result (error) information.
type TopicResult struct {
	// Topic name
	Topic string
	// Error, if any, of result. Check with `Error.Code() != ErrNoError`.
	Error Error
}

// String returns a human-readable representation of a TopicResult.
func (t TopicResult) String() string {
	return ""
}

// TopicSpecification holds parameters for creating a new topic.
// TopicSpecification is analogous to NewTopic in the Java Topic Admin API.
type TopicSpecification struct {
	// Topic name to create.
	Topic string
	// Number of partitions in topic.
	NumPartitions int
	// Default replication factor for the topic's partitions, or zero
	// if an explicit ReplicaAssignment is set.
	ReplicationFactor int
	// (Optional) Explicit replica assignment. The outer array is
	// indexed by the partition number, while the inner per-partition array
	// contains the replica broker ids. The first broker in each
	// broker id list will be the preferred replica.
	ReplicaAssignment [][]int32
	// Topic configuration.
	Config map[string]string
}

// PartitionsSpecification holds parameters for creating additional partitions for a topic.
// PartitionsSpecification is analogous to NewPartitions in the Java Topic Admin API.
type PartitionsSpecification struct {
	// Topic to create more partitions for.
	Topic string
	// New partition count for topic, must be higher than current partition count.
	IncreaseTo int
	// (Optional) Explicit replica assignment. The outer array is
	// indexed by the new partition index (i.e., 0 for the first added
	// partition), while the inner per-partition array
	// contains the replica broker ids. The first broker in each
	// broker id list will be the preferred replica.
	ReplicaAssignment [][]int32
}

// ResourceType represents an Apache Kafka resource type
type ResourceType int

const (
	// ResourceUnknown - Unknown
	ResourceUnknown = ResourceType(0)
	// ResourceAny - match any resource type (DescribeConfigs)
	ResourceAny = ResourceType(0)
	// ResourceTopic - Topic
	ResourceTopic = ResourceType(0)
	// ResourceGroup - Group
	ResourceGroup = ResourceType(0)
	// ResourceBroker - Broker
	ResourceBroker = ResourceType(0)
)

// String returns the human-readable representation of a ResourceType
func (t ResourceType) String() string {
	return ""
}

// ResourceTypeFromString translates a resource type name/string to
// a ResourceType value.
func ResourceTypeFromString(typeString string) (ResourceType, error) {
	return 0, errors.New("")
}

// ConfigSource represents an Apache Kafka config source
type ConfigSource int

const (
	// ConfigSourceUnknown is the default value
	ConfigSourceUnknown = ConfigSource(0)
	// ConfigSourceDynamicTopic is dynamic topic config that is configured for a specific topic
	ConfigSourceDynamicTopic = ConfigSource(0)
	// ConfigSourceDynamicBroker is dynamic broker config that is configured for a specific broker
	ConfigSourceDynamicBroker = ConfigSource(0)
	// ConfigSourceDynamicDefaultBroker is dynamic broker config that is configured as default for all brokers in the cluster
	ConfigSourceDynamicDefaultBroker = ConfigSource(0)
	// ConfigSourceStaticBroker is static broker config provided as broker properties at startup (e.g. from server.properties file)
	ConfigSourceStaticBroker = ConfigSource(0)
	// ConfigSourceDefault is built-in default configuration for configs that have a default value
	ConfigSourceDefault = ConfigSource(0)
)

// String returns the human-readable representation of a ConfigSource type
func (t ConfigSource) String() string {
	return ""
}

// ConfigResource holds parameters for altering an Apache Kafka configuration resource
type ConfigResource struct {
	// Type of resource to set.
	Type ResourceType
	// Name of resource to set.
	Name string
	// Config entries to set.
	// Configuration updates are atomic, any configuration property not provided
	// here will be reverted (by the broker) to its default value.
	// Use DescribeConfigs to retrieve the list of current configuration entry values.
	Config []ConfigEntry
}

// String returns a human-readable representation of a ConfigResource
func (c ConfigResource) String() string {
	return ""
}

// AlterOperation specifies the operation to perform on the ConfigEntry.
// Currently only AlterOperationSet.
type AlterOperation int

const (
	// AlterOperationSet sets/overwrites the configuration setting.
	AlterOperationSet = iota
)

// String returns the human-readable representation of an AlterOperation
func (o AlterOperation) String() string {
	return ""
}

// ConfigEntry holds parameters for altering a resource's configuration.
type ConfigEntry struct {
	// Name of configuration entry, e.g., topic configuration property name.
	Name string
	// Value of configuration entry.
	Value string
	// Operation to perform on the entry.
	Operation AlterOperation
}

// StringMapToConfigEntries creates a new map of ConfigEntry objects from the
// provided string map. The AlterOperation is set on each created entry.
func StringMapToConfigEntries(stringMap map[string]string, operation AlterOperation) []ConfigEntry {
	return nil
}

// String returns a human-readable representation of a ConfigEntry.
func (c ConfigEntry) String() string {
	return ""
}

// ConfigEntryResult contains the result of a single configuration entry from a
// DescribeConfigs request.
type ConfigEntryResult struct {
	// Name of configuration entry, e.g., topic configuration property name.
	Name string
	// Value of configuration entry.
	Value string
	// Source indicates the configuration source.
	Source ConfigSource
	// IsReadOnly indicates whether the configuration entry can be altered.
	IsReadOnly bool
	// IsSensitive indicates whether the configuration entry contains sensitive information, in which case the value will be unset.
	IsSensitive bool
	// IsSynonym indicates whether the configuration entry is a synonym for another configuration property.
	IsSynonym bool
	// Synonyms contains a map of configuration entries that are synonyms to this configuration entry.
	Synonyms map[string]ConfigEntryResult
}

// String returns a human-readable representation of a ConfigEntryResult.
func (c ConfigEntryResult) String() string {
	return ""
}

// ConfigResourceResult provides the result for a resource from a AlterConfigs or
// DescribeConfigs request.
type ConfigResourceResult struct {
	// Type of returned result resource.
	Type ResourceType
	// Name of returned result resource.
	Name string
	// Error, if any, of returned result resource.
	Error Error
	// Config entries, if any, of returned result resource.
	Config map[string]ConfigEntryResult
}

// String returns a human-readable representation of a ConfigResourceResult.
func (c ConfigResourceResult) String() string {
	return ""
}

// ClusterID returns the cluster ID as reported in broker metadata.
//
// Note on cancellation: Although the underlying C function respects the
// timeout, it currently cannot be manually cancelled. That means manually
// cancelling the context will block until the C function call returns.
//
// Requires broker version >= 0.10.0.
func (a *AdminClient) ClusterID(ctx context.Context) (clusterID string, err error) {
	return "", errors.New("")
}

// ControllerID returns the broker ID of the current controller as reported in
// broker metadata.
//
// Note on cancellation: Although the underlying C function respects the
// timeout, it currently cannot be manually cancelled. That means manually
// cancelling the context will block until the C function call returns.
//
// Requires broker version >= 0.10.0.
func (a *AdminClient) ControllerID(ctx context.Context) (controllerID int32, err error) {
	return 0, errors.New("")
}

// CreateTopics creates topics in cluster.
//
// The list of TopicSpecification objects define the per-topic partition count, replicas, etc.
//
// Topic creation is non-atomic and may succeed for some topics but fail for others,
// make sure to check the result for topic-specific errors.
//
// Note: TopicSpecification is analogous to NewTopic in the Java Topic Admin API.
func (a *AdminClient) CreateTopics(ctx context.Context, topics []TopicSpecification, options ...CreateTopicsAdminOption) (result []TopicResult, err error) {
	return nil, errors.New("")
}

// DeleteTopics deletes a batch of topics.
//
// This operation is not transactional and may succeed for a subset of topics while
// failing others.
// It may take several seconds after the DeleteTopics result returns success for
// all the brokers to become aware that the topics are gone. During this time,
// topic metadata and configuration may continue to return information about deleted topics.
//
// Requires broker version >= 0.10.1.0
func (a *AdminClient) DeleteTopics(ctx context.Context, topics []string, options ...DeleteTopicsAdminOption) (result []TopicResult, err error) {
	return nil, errors.New("")
}

// CreatePartitions creates additional partitions for topics.
func (a *AdminClient) CreatePartitions(ctx context.Context, partitions []PartitionsSpecification, options ...CreatePartitionsAdminOption) (result []TopicResult, err error) {
	return nil, errors.New("")
}

// AlterConfigs alters/updates cluster resource configuration.
//
// Updates are not transactional so they may succeed for a subset
// of the provided resources while others fail.
// The configuration for a particular resource is updated atomically,
// replacing values using the provided ConfigEntrys and reverting
// unspecified ConfigEntrys to their default values.
//
// Requires broker version >=0.11.0.0
//
// AlterConfigs will replace all existing configuration for
// the provided resources with the new configuration given,
// reverting all other configuration to their default values.
//
// Multiple resources and resource types may be set, but at most one
// resource of type ResourceBroker is allowed per call since these
// resource requests must be sent to the broker specified in the resource.
func (a *AdminClient) AlterConfigs(ctx context.Context, resources []ConfigResource, options ...AlterConfigsAdminOption) (result []ConfigResourceResult, err error) {
	return nil, errors.New("")
}

// DescribeConfigs retrieves configuration for cluster resources.
//
// The returned configuration includes default values, use
// ConfigEntryResult.IsDefault or ConfigEntryResult.Source to distinguish
// default values from manually configured settings.
//
// The value of config entries where .IsSensitive is true
// will always be nil to avoid disclosing sensitive
// information, such as security settings.
//
// Configuration entries where .IsReadOnly is true can't be modified
// (with AlterConfigs).
//
// Synonym configuration entries are returned if the broker supports
// it (broker version >= 1.1.0). See .Synonyms.
//
// Requires broker version >=0.11.0.0
//
// Multiple resources and resource types may be requested, but at most
// one resource of type ResourceBroker is allowed per call
// since these resource requests must be sent to the broker specified
// in the resource.
func (a *AdminClient) DescribeConfigs(ctx context.Context, resources []ConfigResource, options ...DescribeConfigsAdminOption) (result []ConfigResourceResult, err error) {
	return nil, errors.New("")
}

// GetMetadata queries broker for cluster and topic metadata.
// If topic is non-nil only information about that topic is returned, else if
// allTopics is false only information about locally used topics is returned,
// else information about all topics is returned.
// GetMetadata is equivalent to listTopics, describeTopics and describeCluster in the Java API.
func (a *AdminClient) GetMetadata(topic *string, allTopics bool, timeoutMs int) (*Metadata, error) {
	return nil, errors.New("")
}

// String returns a human readable name for an AdminClient instance
func (a *AdminClient) String() string {
	return ""
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
func (a *AdminClient) SetOAuthBearerToken(oauthBearerToken OAuthBearerToken) error {
	return errors.New("")
}

// SetOAuthBearerTokenFailure sets the error message describing why token
// retrieval/setting failed; it also schedules a new token refresh event for 10
// seconds later so the attempt may be retried. It will return nil on
// success, otherwise an error if:
// 1) SASL/OAUTHBEARER is not supported by the underlying librdkafka build;
// 2) SASL/OAUTHBEARER is supported but is not configured as the client's
// authentication mechanism.
func (a *AdminClient) SetOAuthBearerTokenFailure(errstr string) error {
	return errors.New("")
}

// Close an AdminClient instance.
func (a *AdminClient) Close() {
}

// NewAdminClient creats a new AdminClient instance with a new underlying client instance
func NewAdminClient(conf *ConfigMap) (*AdminClient, error) {
	return nil, errors.New("")
}

// NewAdminClientFromProducer derives a new AdminClient from an existing Producer instance.
// The AdminClient will use the same configuration and connections as the parent instance.
func NewAdminClientFromProducer(p *Producer) (a *AdminClient, err error) {
	return nil, errors.New("")
}

// NewAdminClientFromConsumer derives a new AdminClient from an existing Consumer instance.
// The AdminClient will use the same configuration and connections as the parent instance.
func NewAdminClientFromConsumer(c *Consumer) (a *AdminClient, err error) {
	return nil, errors.New("")
}
