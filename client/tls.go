package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net"
	"time"
)

// TLSClient holds TLS client connection.
type TLSClient struct {
	conn      net.Conn
	tlsConfig *tls.Config
}

// NewTLSClient returns TLSClient handle.
func NewTLSClient(ctx context.Context, addr string, clientCert []byte, clientCertKey []byte, caCert []byte, serverName string) (*TLSClient, error) {

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	cert, err := tls.X509KeyPair(clientCert, clientCertKey)
	if err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
		MinVersion:   tls.VersionTLS12,
		MaxVersion:   tls.VersionTLS12,
		ServerName:   serverName,
	}

	d := tls.Dialer{Config: tlsConfig}
	conn, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}

	return &TLSClient{conn: conn, tlsConfig: tlsConfig}, nil
}

// Send sends data on the wire.
func (t *TLSClient) Send(data string) error {

	if err := writeData(t.conn, data); err == nil {
		return nil
	}

	var d net.Dialer
	d.Timeout = defaultTimeout
	conn, err := tls.DialWithDialer(&d, t.conn.RemoteAddr().Network(), t.conn.RemoteAddr().String(), t.tlsConfig)
	if err != nil {
		return err
	}

	t.conn.Close() // nolint:errcheck

	t.conn = conn

	return writeData(t.conn, data)
}

// Protocol returns TLS.
func (t *TLSClient) Protocol() Protocol {
	return TLS
}

// SetTimeout sets connection timeout.
func (t *TLSClient) SetTimeout(dur time.Duration) error {
	return t.conn.SetDeadline(time.Now().Add(dur))
}

// Close closes the connection.
func (t *TLSClient) Close() error {
	return t.conn.Close()
}
