package client

import (
	"context"
	"net"
	"time"
)

// UnixClient holds Unix connection.
type UnixClient struct {
	conn net.Conn
}

// NewUnixClient return Unix client.
func NewUnixClient(ctx context.Context) (*UnixClient, error) {

	var conn net.Conn
	var err error

	// Ref https://github.com/gabriel-samfira/syslog/blob/969ede9b50d934db936f618897d1b1a64dd4edf1/syslog_unix.go#L18
L:
	for _, network := range []string{"unixgram", "unix"} {
		for _, path := range []string{"/dev/log", "/var/run/syslog"} {

			var d net.Dialer
			conn, err = d.DialContext(ctx, network, path)
			if err != nil {
				continue
			} else {
				break L
			}
		}
	}

	if err != nil {
		return nil, err
	}

	return &UnixClient{conn: conn}, nil
}

// Send sends data on wire.
func (t *UnixClient) Send(data string) error {
	return writeData(t.conn, data)
}

// Protocol returns Unix.
func (t *UnixClient) Protocol() Protocol {
	return Unix
}

// SetTimeout sets connection timeout.
func (t *UnixClient) SetTimeout(dur time.Duration) error {
	return t.conn.SetDeadline(time.Now().Add(dur))
}

// Close closes the connection.
func (t *UnixClient) Close() error {
	return t.conn.Close()
}
