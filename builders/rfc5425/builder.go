package rfc5425

import (
	"fmt"
	"time"

	"github.com/sibicramesh/syslog-go/builders/rfc5424"
)

// RFC https://datatracker.ietf.org/doc/html/rfc5425

// Builder holds rfc5425 syslog message fields.
type Builder struct {
	rfc5424 *rfc5424.Builder
}

// New returns rfc5425 builder.
func New() *Builder {
	return &Builder{
		rfc5424: rfc5424.New(),
	}
}

// SetPriority sets priority.
func (r *Builder) SetPriority(facility int, severity int) error {
	return r.rfc5424.SetPriority(facility, severity)
}

// SetTimestamp sets timestamp.
func (r *Builder) SetTimestamp(stamp time.Time) error {
	return r.rfc5424.SetTimestamp(stamp)
}

// SetHostname sets hostname.
func (r *Builder) SetHostname(hostname string) error {
	return r.rfc5424.SetHostname(hostname)
}

// SetMsg sets content.
func (r *Builder) SetMsg(_ string, content string) error {
	return r.rfc5424.SetMsg("", content)
}

// SetAppName is sets application/process name.
func (r *Builder) SetAppName(appName string) error {
	return r.rfc5424.SetAppName(appName)
}

// String returns valid syslog message.
// MSGLEN MSG
func (r *Builder) String() string {

	rfc5424Msg := r.rfc5424.String()

	return fmt.Sprintf("%d %s", len(rfc5424Msg), rfc5424Msg)
}
