package client

import (
	"errors"
	"time"
)

// FakeClient is used for testing.
type FakeClient struct {
	err bool
}

// NewFakeClient returns FakeClient handle.
func NewFakeClient(err bool) *FakeClient {
	return &FakeClient{err: err}
}

// Send is a noop method.
func (f *FakeClient) Send(_ string) error {

	if f.err {
		return errors.New("fake error")
	}

	return nil
}

// Protocol returns TCP.
func (f *FakeClient) Protocol() Protocol {
	return TCP
}

// SetTimeout sets connection timeout.
func (f *FakeClient) SetTimeout(dur time.Duration) error {
	return nil
}

// Close is a noop method.
func (f *FakeClient) Close() error {
	return nil
}
