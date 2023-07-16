package client

import (
	"context"
	"net"
	"time"
)

// UDPClient holds UDP client connection.
type UDPClient struct {
	conn net.Conn
}

// NewUDPClient returns UDP client.
func NewUDPClient(ctx context.Context, addr string) (*UDPClient, error) {

	var d net.Dialer
	conn, err := d.DialContext(ctx, "udp", addr)
	if err != nil {
		return nil, err
	}

	return &UDPClient{conn: conn}, nil
}

// Send sends data on the wire.
func (t *UDPClient) Send(data string) error {

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

// Protocol returns UDP.
func (t *UDPClient) Protocol() Protocol {
	return UDP
}

// SetTimeout sets connection timeout.
func (t *UDPClient) SetTimeout(dur time.Duration) error {
	return t.conn.SetDeadline(time.Now().Add(dur))
}

// Close closes the connection.
func (t *UDPClient) Close() error {
	return t.conn.Close()
}
