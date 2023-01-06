package kafka

import (
	"errors"
	"sync"
	"time"
)

type OAuthBearerToken struct {
	TokenValue string
	Expiration time.Time
	Principal  string
	Extensions map[string]string
}

type Handle interface {
	SetOAuthBearerToken(oauthBearerToken OAuthBearerToken) error

	SetOAuthBearerTokenFailure(errstr string) error
	gethandle() *handle
}

type handle struct {
	// Forward logs from librdkafka log queue to logs channel.
	logs          chan LogEvent
	closeLogsChan bool

	// Topic <-> rkt caches
	rktCacheLock sync.Mutex

	// Cached instance name to avoid CGo call in String()
	name string

	//
	// cgo map
	// Maps C callbacks based on cgoid back to its Go object
	cgoLock   sync.Mutex
	cgoidNext uintptr
	cgomap    map[int]cgoif

	//
	// producer
	//
	p *Producer

	// Forward delivery reports on Producer.Events channel
	fwdDr bool

	// Enabled message fields for delivery reports and consumed messages.
	msgFields *messageFields

	//
	// consumer
	//
	c *Consumer

	// WaitGroup to wait for spawned go-routines to finish.
	waitGroup sync.WaitGroup
}

func (h *handle) String() string {
	return h.name
}

type cgoif interface{}

func (h *handle) setOAuthBearerToken(oauthBearerToken OAuthBearerToken) error {
	return errors.New("unimplemented")
}

func (h *handle) setOAuthBearerTokenFailure(errstr string) error {
	return errors.New("unimplemented")
}

type messageFields struct {
	Key     bool
	Value   bool
	Headers bool
}
