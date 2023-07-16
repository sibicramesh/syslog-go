package rfc5424

import (
	"fmt"
	"time"

	"github.com/sibicramesh/syslog-go/builders/common"
)

// RFC https://datatracker.ietf.org/doc/html/rfc5424

// Builder holds rfc5424 syslog message fields.
type Builder struct {
	priority  int
	timestamp string
	hostname  string
	appName   string
	msg       string
}

// New returns rfc5424 builder.
func New() *Builder {
	return &Builder{
		priority:  13,
		timestamp: "-",
		hostname:  "-",
		appName:   "-",
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

	r.timestamp = stamp.Format(time.RFC3339)

	return nil
}

// SetHostname sets hostname.
func (r *Builder) SetHostname(hostname string) error {

	r.hostname = hostname

	return nil
}

// SetMsg sets content.
func (r *Builder) SetMsg(_ string, content string) error {

	r.msg = fmt.Sprintf("\xEF\xBB\xBF %s", content)

	return nil
}

// SetAppName is sets application/process name.
func (r *Builder) SetAppName(appName string) error {

	r.appName = appName

	return nil
}

// String returns valid syslog message.
// <PRI>TIMESTAMP HOSTNAME APPNAME - - - MSG
func (r *Builder) String() string {
	return fmt.Sprintf("<%d>1 %s %s %s - - - %s\n", r.priority, r.timestamp, r.hostname, r.appName, r.msg)
}
