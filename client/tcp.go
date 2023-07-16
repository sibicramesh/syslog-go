package client

import (
	"context"
	"net"
	"time"
)

// TCPClient holds TCP client connection.
type TCPClient struct {
	conn net.Conn
}

// NewTCPClient returns TCPClient handle.
func NewTCPClient(ctx context.Context, addr string) (*TCPClient, error) {

	var d net.Dialer
	conn, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}

	return &TCPClient{conn: conn}, nil
}

// Send sends the data on the wire.
func (t *TCPClient) Send(data string) error {

	if err := writeData(t.conn, data); err == nil {
		return nil
	}

	conn, err := net.DialTimeout(t.conn.RemoteAddr().Network(), t.conn.RemoteAddr().String(), defaultTimeout)
	if err != nil {
		return err
	}

	t.conn.Close() // nolint:errcheck

	t.conn = conn

	return writeData(t.conn, data)
}

// Protocol returns TCP.
func (t *TCPClient) Protocol() Protocol {
	return TCP
}

// SetTimeout sets connection timeout.
func (t *TCPClient) SetTimeout(dur time.Duration) error {
	return t.conn.SetDeadline(time.Now().Add(dur))
}

// Close closes the connection.
func (t *TCPClient) Close() error {
	return t.conn.Close()
}
