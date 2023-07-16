// nolint
package rfc5424

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_BuildRFC5424(t *testing.T) {

	// valid message
	rfc5424Builder := New()
	rfc5424Builder.SetPriority(1, 5)
	tim, err := time.Parse(time.Stamp, "Nov 4 16:59:58")
	require.Nil(t, err)
	rfc5424Builder.SetTimestamp(tim)
	rfc5424Builder.SetHostname("my.host.name")
	rfc5424Builder.SetMsg("mytest", "just a random message")
	rfc5424Builder.SetAppName("myapp")
	require.Equal(t, "<13>1 0000-11-04T16:59:58Z my.host.name myapp - - - \ufeff just a random message\n", rfc5424Builder.String())

	// invalid facility
	rfc5424Builder = New()
	err = rfc5424Builder.SetPriority(24, 5)
	require.NotNil(t, err)

	// invalid severity
	rfc5424Builder = New()
	err = rfc5424Builder.SetPriority(1, 8)
	require.NotNil(t, err)
}
