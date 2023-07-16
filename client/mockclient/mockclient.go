// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sibicramesh/syslog-go (interfaces: Client)

// Package mockclient is a generated GoMock package.
package mockclient

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	client "github.com/sibicramesh/syslog-go/client"
)

// MockClient is a mock of Client interface.
// nolint
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
// nolint
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
// nolint
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
// nolint
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
// nolint
func (m *MockClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
// nolint
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// Protocol mocks base method.
// nolint
func (m *MockClient) Protocol() client.Protocol {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Protocol")
	ret0, _ := ret[0].(client.Protocol)
	return ret0
}

// Protocol indicates an expected call of Protocol.
// nolint
func (mr *MockClientMockRecorder) Protocol() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Protocol", reflect.TypeOf((*MockClient)(nil).Protocol))
}

// Send mocks base method.
// nolint
func (m *MockClient) Send(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send.
// nolint
func (mr *MockClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockClient)(nil).Send), arg0)
}

// SetTimeout mocks base method.
// nolint
func (m *MockClient) SetTimeout(arg0 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTimeout", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTimeout indicates an expected call of SetTimeout.
// nolint
func (mr *MockClientMockRecorder) SetTimeout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimeout", reflect.TypeOf((*MockClient)(nil).SetTimeout), arg0)
}