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
)

// ConfigValue supports the following types:
//
//	bool, int, string, any type with the standard String() interface
type ConfigValue interface{}

// ConfigMap is a map containing standard librdkafka configuration properties as documented in:
// https://github.com/edenhill/librdkafka/tree/master/CONFIGURATION.md
//
// The special property "default.topic.config" (optional) is a ConfigMap
// containing default topic configuration properties.
//
// The use of "default.topic.config" is deprecated,
// topic configuration properties shall be specified in the standard ConfigMap.
// For backwards compatibility, "default.topic.config" (if supplied)
// takes precedence.
type ConfigMap map[string]ConfigValue

// SetKey sets configuration property key to value.
//
// For user convenience a key prefixed with {topic}. will be
// set on the "default.topic.config" sub-map, this use is deprecated.
func (m ConfigMap) SetKey(key string, value ConfigValue) error {
	return errors.New("")
}

// Set implements flag.Set (command line argument parser) as a convenience
// for `-X key=value` config.
func (m ConfigMap) Set(kv string) error {
	return errors.New("")
}

// Get finds the given key in the ConfigMap and returns its value.
// If the key is not found `defval` is returned.
// If the key is found but the type does not match that of `defval` (unless nil)
// an ErrInvalidArg error is returned.
func (m ConfigMap) Get(key string, defval ConfigValue) (ConfigValue, error) {
	return 0, errors.New("")
}
