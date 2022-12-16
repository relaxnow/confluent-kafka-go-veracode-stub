package kafka

// ConfigValue supports the following types:
//
//	bool, int, string, any type with the standard String() interface
type ConfigValue interface{}

type ConfigMap map[string]ConfigValue
