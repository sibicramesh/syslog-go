// nolint
package rfc3164

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_BuildRFC3164(t *testing.T) {

	// valid message
	rfc3164Builder := New()
	rfc3164Builder.SetPriority(1, 5)
	tim, err := time.Parse(time.Stamp, "Nov 4 16:59:58")
	require.Nil(t, err)
	rfc3164Builder.SetTimestamp(tim)
	rfc3164Builder.SetHostname("my.host.name")
	rfc3164Builder.SetMsg("mytest", "just a random message")
	require.Equal(t, "<13>Nov  4 16:59:58 my.host.name mytest: just a random message", rfc3164Builder.String())

	// invalid facility
	rfc3164Builder = New()
	err = rfc3164Builder.SetPriority(24, 5)
	require.NotNil(t, err)

	// invalid severity
	rfc3164Builder = New()
	err = rfc3164Builder.SetPriority(1, 8)
	require.NotNil(t, err)

	// invalid hostname
	rfc3164Builder = New()
	err = rfc3164Builder.SetHostname("my host")
	require.NotNil(t, err)

	// invalid tag
	rfc3164Builder = New()
	err = rfc3164Builder.SetMsg("asdfghlokjhgerfghjksdngthedrftgss", "my host")
	require.NotNil(t, err)

	// missing timestamp
	rfc3164Builder = New()
	require.NotEmpty(t, rfc3164Builder.String())
	require.NotEmpty(t, rfc3164Builder.timestamp)
}
