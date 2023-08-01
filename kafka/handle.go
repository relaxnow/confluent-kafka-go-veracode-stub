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

// OAuthBearerToken represents the data to be transmitted
// to a broker during SASL/OAUTHBEARER authentication.
type OAuthBearerToken struct {
	// Token value, often (but not necessarily) a JWS compact serialization
	// as per https://tools.ietf.org/html/rfc7515#section-3.1; it must meet
	// the regular expression for a SASL/OAUTHBEARER value defined at
	// https://tools.ietf.org/html/rfc7628#section-3.1
	TokenValue string
	// Metadata about the token indicating when it expires (local time);
	// it must represent a time in the future
	Expiration time.Time
	// Metadata about the token indicating the Kafka principal name
	// to which it applies (for example, "admin")
	Principal string
	// SASL extensions, if any, to be communicated to the broker during
	// authentication (all keys and values of which must meet the regular
	// expressions defined at https://tools.ietf.org/html/rfc7628#section-3.1,
	// and it must not contain the reserved "auth" key)
	Extensions map[string]string
}

// Handle represents a generic client handle containing common parts for
// both Producer and Consumer.
type Handle interface {
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
	SetOAuthBearerToken(oauthBearerToken OAuthBearerToken) error

	// SetOAuthBearerTokenFailure sets the error message describing why token
	// retrieval/setting failed; it also schedules a new token refresh event for 10
	// seconds later so the attempt may be retried. It will return nil on
	// success, otherwise an error if:
	// 1) SASL/OAUTHBEARER is not supported by the underlying librdkafka build;
	// 2) SASL/OAUTHBEARER is supported but is not configured as the client's
	// authentication mechanism.
	SetOAuthBearerTokenFailure(errstr string) error
	gethandle() *handle
}

type handle struct {
	// Forward logs from librdkafka log queue to logs channel.
}

func (h *handle) String() string {
	return ""
}
