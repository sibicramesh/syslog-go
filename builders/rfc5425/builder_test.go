// nolint
package rfc5425

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_BuildRFC35425(t *testing.T) {

	// valid message
	rfc5Builder := New()
	rfc5Builder.SetPriority(1, 5)
	tim, err := time.Parse(time.Stamp, "Nov 4 16:59:58")
	require.Nil(t, err)
	rfc5Builder.SetTimestamp(tim)
	rfc5Builder.SetHostname("my.host.name")
	rfc5Builder.SetMsg("mytest", "just a random message")
	rfc5Builder.SetAppName("myapp")
	require.Equal(t, "78 <13>1 0000-11-04T16:59:58Z my.host.name myapp - - - \ufeff just a random message\n", rfc5Builder.String())

	// invalid facility
	rfc5Builder = New()
	err = rfc5Builder.SetPriority(24, 5)
	require.NotNil(t, err)

	// invalid severity
	rfc5Builder = New()
	err = rfc5Builder.SetPriority(1, 8)
	require.NotNil(t, err)
}
