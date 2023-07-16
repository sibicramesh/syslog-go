//go:build !linux

package client

import (
	"context"
	"time"
)

// UnixClient holds Unix connection.
type UnixClient struct {
}

// NewUnixClient return Unix client.
func NewUnixClient(ctx context.Context) (*UnixClient, error) {
	return &UnixClient{}, nil
}

// Send sends data on wire.
func (t *UnixClient) Send(data string) error {
	return nil
}

// Protocol returns Unix.
func (t *UnixClient) Protocol() Protocol {
	return Unix
}

// SetTimeout sets connection timeout.
func (t *UnixClient) SetTimeout(_ time.Duration) error {
	return nil
}

// Close closes the connection.
func (t *UnixClient) Close() error {
	return nil
}
