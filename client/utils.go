package client

import (
	"fmt"
	"net"
	"time"
)

// Protocol types.
type Protocol int

// Different protocols.
const (
	UDP = iota
	TCP
	TLS
	Unix
)

const defaultTimeout = 15 * time.Second

// Protocol returns string representation of protocol.
func (p Protocol) String() string {

	switch p {
	case UDP:
		return "UDP"
	case TCP:
		return "TCP"
	case TLS:
		return "TLS"
	case Unix:
		return "Unix"
	default:
		return "UDP"
	}
}

func writeData(conn net.Conn, data string) error {

	n, err := fmt.Fprint(conn, data)
	if err != nil {
		return fmt.Errorf("unable to write syslog message: %w", err)
	}

	if n != len(data) {
		return fmt.Errorf("unable to write complete syslog message: %d", n)
	}

	return nil
}
