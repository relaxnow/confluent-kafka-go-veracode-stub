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

// Error provides a Kafka-specific error container
type Error struct {
	code             ErrorCode
	str              string
	fatal            bool
	retriable        bool
	txnRequiresAbort bool
}

// NewError creates a new Error.
func NewError(code ErrorCode, str string, fatal bool) (err Error) {
	return Error{}
}

// Error returns a human readable representation of an Error
// Same as Error.String()
func (e Error) Error() string {
	return ""
}

// String returns a human readable representation of an Error
func (e Error) String() string {
	return ""
}

// Code returns the ErrorCode of an Error
func (e Error) Code() ErrorCode {
	return 0
}

// IsFatal returns true if the error is a fatal error.
// A fatal error indicates the client instance is no longer operable and
// should be terminated. Typical causes include non-recoverable
// idempotent producer errors.
func (e Error) IsFatal() bool {
	return false
}

// IsRetriable returns true if the operation that caused this error
// may be retried.
// This flag is currently only set by the Transactional producer API.
func (e Error) IsRetriable() bool {
	return false
}

// TxnRequiresAbort returns true if the error is an abortable transaction error
// that requires the application to abort the current transaction with
// AbortTransaction() and start a new transaction with BeginTransaction()
// if it wishes to proceed with transactional operations.
// This flag is only set by the Transactional producer API.
func (e Error) TxnRequiresAbort() bool {
	return false
}
