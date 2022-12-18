package kafka

import "time"

// LogEvent represent the log from librdkafka internal log queue
type LogEvent struct {
	Name      string    // Name of client instance
	Tag       string    // Log tag that provides context to the log Message (e.g., "METADATA" or "GRPCOORD")
	Message   string    // Log message
	Level     int       // Log syslog level, lower is more critical.
	Timestamp time.Time // Log timestamp
}
