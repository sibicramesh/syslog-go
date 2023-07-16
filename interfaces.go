package syslog

import (
	"time"

	"github.com/sibicramesh/syslog-go/client"
)

// Builder interface contains the methods required to build
// a valid syslog message. Note that the interface might have
// methods specific to certain syslog RFC's.
type Builder interface {
	SetPriority(facility int, severity int) error
	SetTimestamp(stamp time.Time) error
	SetHostname(hostname string) error
	// tag is only implemented in rfc3164. It should have
	// process or application that generated the message.
	SetMsg(tag string, content string) error
	// Unimplemented in rfc3164, instead one can add the app
	// name in the tag section of the msg.
	SetAppName(appName string) error
	String() string
}

// Client interface contains methods to send a syslog messages.
// It can be a file writer or socket writer as long as a single
// write is performed.
type Client interface {
	Send(data string) error
	Protocol() client.Protocol
	SetTimeout(dur time.Duration) error
	Close() error
}

// Syslog is the main interface which is used to send the
// syslog message in the given format. The obj can be of any
// arbitrary type. Encoding the obj totally depends on the
// implementation and the use case. Complete RFC compliant.
type Syslog interface {
	Send(format Format, facility int, severity int, obj interface{}) error
}
