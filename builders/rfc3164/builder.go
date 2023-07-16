package rfc3164

import (
	"fmt"
	"strings"
	"time"

	"github.com/sibicramesh/syslog-go/builders/common"
)

// RFC https://datatracker.ietf.org/doc/html/rfc3164

// Builder holds rfc3164 syslog message fields.
type Builder struct {
	priority  int
	timestamp string
	hostname  string
	msg       string
}

// New returns rfc3164 builder.
func New() *Builder {
	return &Builder{
		priority: 13,
		hostname: "default",
	}
}

// SetPriority sets priority.
func (r *Builder) SetPriority(facility int, severity int) error {

	priority, err := common.CalculatePriority(facility, severity)
	if err != nil {
		return err
	}

	r.priority = priority

	return nil
}

// SetTimestamp sets timestamp.
func (r *Builder) SetTimestamp(stamp time.Time) error {

	r.timestamp = stamp.Format(time.Stamp)

	return nil
}

// SetHostname sets hostname.
func (r *Builder) SetHostname(hostname string) error {

	if strings.ContainsAny(hostname, " ") {
		return fmt.Errorf("hostname cannot contain space(s): %s", hostname)
	}

	r.hostname = hostname

	return nil
}

// SetMsg sets tag and content.
func (r *Builder) SetMsg(tag string, content string) error {

	if len(tag) > 32 {
		return fmt.Errorf("tag length cannot exceed 32 characters: %s", tag)
	}

	r.msg = tag + ": " + content

	return nil
}

// SetAppName is unimplemented.
func (r *Builder) SetAppName(appName string) error {
	return nil
}

// String returns valid syslog message.
// <PRI>TIMESTAMP HOSTNAME TAG: MSG
func (r *Builder) String() string {

	if r.timestamp == "" {
		r.timestamp = time.Now().Format(time.Stamp)
	}

	return fmt.Sprintf("<%d>%s %s %s", r.priority, r.timestamp, r.hostname, r.msg)
}
