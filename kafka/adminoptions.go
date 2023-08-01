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
	"time"
)

// AdminOptionOperationTimeout sets the broker's operation timeout, such as the
// timeout for CreateTopics to complete the creation of topics on the controller
// before returning a result to the application.
//
// CreateTopics, DeleteTopics, CreatePartitions:
// a value 0 will return immediately after triggering topic
// creation, while > 0 will wait this long for topic creation to propagate
// in cluster.
//
// Default: 0 (return immediately).
//
// Valid for CreateTopics, DeleteTopics, CreatePartitions.
type AdminOptionOperationTimeout struct {
	isSet bool
	val   time.Duration
}

// SetAdminOperationTimeout sets the broker's operation timeout, such as the
// timeout for CreateTopics to complete the creation of topics on the controller
// before returning a result to the application.
//
// CreateTopics, DeleteTopics, CreatePartitions:
// a value 0 will return immediately after triggering topic
// creation, while > 0 will wait this long for topic creation to propagate
// in cluster.
//
// Default: 0 (return immediately).
//
// Valid for CreateTopics, DeleteTopics, CreatePartitions.
func SetAdminOperationTimeout(t time.Duration) (ao AdminOptionOperationTimeout) {
	return AdminOptionOperationTimeout{}
}

// AdminOptionRequestTimeout sets the overall request timeout, including broker
// lookup, request transmission, operation time on broker, and response.
//
// Default: `socket.timeout.ms`.
//
// Valid for all Admin API methods.
type AdminOptionRequestTimeout struct {
	isSet bool
	val   time.Duration
}

// SetAdminRequestTimeout sets the overall request timeout, including broker
// lookup, request transmission, operation time on broker, and response.
//
// Default: `socket.timeout.ms`.
//
// Valid for all Admin API methods.
func SetAdminRequestTimeout(t time.Duration) (ao AdminOptionRequestTimeout) {
	return AdminOptionRequestTimeout{}
}

// AdminOptionValidateOnly tells the broker to only validate the request,
// without performing the requested operation (create topics, etc).
//
// Default: false.
//
// Valid for CreateTopics, CreatePartitions, AlterConfigs
type AdminOptionValidateOnly struct {
	isSet bool
	val   bool
}

// SetAdminValidateOnly tells the broker to only validate the request,
// without performing the requested operation (create topics, etc).
//
// Default: false.
//
// Valid for CreateTopics, DeleteTopics, CreatePartitions, AlterConfigs
func SetAdminValidateOnly(validateOnly bool) (ao AdminOptionValidateOnly) {
	return AdminOptionValidateOnly{}
}

// CreateTopicsAdminOption - see setters.
//
// See SetAdminRequestTimeout, SetAdminOperationTimeout, SetAdminValidateOnly.
type CreateTopicsAdminOption interface {
	supportsCreateTopics()
	apply(cOptions int) error
}

// DeleteTopicsAdminOption - see setters.
//
// See SetAdminRequestTimeout, SetAdminOperationTimeout.
type DeleteTopicsAdminOption interface {
	supportsDeleteTopics()
	apply(cOptions int) error
}

// CreatePartitionsAdminOption - see setters.
//
// See SetAdminRequestTimeout, SetAdminOperationTimeout, SetAdminValidateOnly.
type CreatePartitionsAdminOption interface {
	supportsCreatePartitions()
	apply(cOptions int) error
}

// AlterConfigsAdminOption - see setters.
//
// See SetAdminRequestTimeout, SetAdminValidateOnly, SetAdminIncremental.
type AlterConfigsAdminOption interface {
	supportsAlterConfigs()
	apply(cOptions int) error
}

// DescribeConfigsAdminOption - see setters.
//
// See SetAdminRequestTimeout.
type DescribeConfigsAdminOption interface {
	supportsDescribeConfigs()
	apply(cOptions int) error
}

// AdminOption is a generic type not to be used directly.
//
// See CreateTopicsAdminOption et.al.
type AdminOption interface {
	apply(cOptions int) error
}
