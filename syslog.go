package syslog

import (
	"fmt"
	"os"
	"time"

	"github.com/mitchellh/go-ps"
	"github.com/sibicramesh/syslog-go/builders/rfc3164"
	"github.com/sibicramesh/syslog-go/builders/rfc5424"
	"github.com/sibicramesh/syslog-go/builders/rfc5425"
	syslogClient "github.com/sibicramesh/syslog-go/client"
	"go.aporeto.io/elemental"
)

// Format represents the different syslog formats to choose from.
type Format int

// Different formats.
const (
	Auto Format = iota
	RFC3164
	RFC5424
	RFC5425
)

// String returns string representation of format.
func (f Format) String() string {

	switch f {
	case Auto:
		return "Auto"
	case RFC3164:
		return "RFC3164"
	case RFC5424:
		return "RFC5424"
	case RFC5425:
		return "RFC5425"
	default:
		return "Auto"
	}
}

type syslog struct {
	client Client

	rfc3164 *rfc3164.Builder
	rfc5424 *rfc5424.Builder
	rfc5425 *rfc5425.Builder

	hostname    string
	processName string
}

// New returns a snew Syslog handle to dispatch syslog messages.
func New(client Client) (Syslog, error) {

	var hostname string

	// If the server is local, the system already writes the hostname
	// in the syslog message, no need to redo it.
	if client.Protocol() != syslogClient.Unix {
		var err error
		hostname, err = os.Hostname()
		if err != nil {
			return nil, fmt.Errorf("unable to get hostname: %w", err)
		}
	}

	process, err := ps.FindProcess(os.Getpid())
	if err != nil {
		return nil, fmt.Errorf("unable to get process info: %w", err)
	}

	return &syslog{
		client:      client,
		rfc3164:     rfc3164.New(),
		rfc5424:     rfc5424.New(),
		rfc5425:     rfc5425.New(),
		hostname:    hostname,
		processName: process.Executable(),
	}, nil
}

// Send dispatches a single syslog messages in the given format
// and returns the error. It encodes the obj in JSON format.
func (s *syslog) Send(format Format, facility int, severity int, obj interface{}) error {

	content, err := elemental.Encode(elemental.EncodingTypeJSON, obj)
	if err != nil {
		return fmt.Errorf("unable to encode object to JSON: %w", err)
	}

	builder := s.getBuilder(format)

	msg, err := s.build(builder, facility, severity, content)
	if err != nil {
		return err
	}

	return s.client.Send(msg)
}

func (s *syslog) getBuilder(format Format) Builder {

	var builder Builder

	switch format {
	case Auto:

		switch s.client.Protocol() {
		case syslogClient.UDP:
			builder = s.rfc3164

		case syslogClient.TCP:
			builder = s.rfc5424

		case syslogClient.TLS:
			builder = s.rfc5425
		}

	case RFC3164:
		builder = s.rfc3164

	case RFC5424:
		builder = s.rfc5424

	case RFC5425:
		builder = s.rfc5425
	}

	if builder == nil {
		return s.rfc3164
	}

	return builder
}

func (s *syslog) build(b Builder, facility int, severity int, content []byte) (string, error) {

	if err := b.SetPriority(facility, severity); err != nil {
		return "", err
	}

	if err := b.SetTimestamp(time.Now()); err != nil {
		return "", err
	}

	if err := b.SetHostname(s.hostname); err != nil {
		return "", err
	}

	if err := b.SetMsg(s.processName, string(content)); err != nil {
		return "", err
	}

	if err := b.SetAppName(s.processName); err != nil {
		return "", err
	}

	return b.String(), nil
}
