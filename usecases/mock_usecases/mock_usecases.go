// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/int128/ghcp/usecases/interfaces (interfaces: Push)

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	interfaces "github.com/int128/ghcp/usecases/interfaces"
	reflect "reflect"
)

// MockPush is a mock of Push interface
type MockPush struct {
	ctrl     *gomock.Controller
	recorder *MockPushMockRecorder
}

// MockPushMockRecorder is the mock recorder for MockPush
type MockPushMockRecorder struct {
	mock *MockPush
}

// NewMockPush creates a new mock instance
func NewMockPush(ctrl *gomock.Controller) *MockPush {
	mock := &MockPush{ctrl: ctrl}
	mock.recorder = &MockPushMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPush) EXPECT() *MockPushMockRecorder {
	return m.recorder
}

// Do mocks base method
func (m *MockPush) Do(arg0 context.Context, arg1 interfaces.PushIn) error {
	ret := m.ctrl.Call(m, "Do", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do
func (mr *MockPushMockRecorder) Do(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockPush)(nil).Do), arg0, arg1)
}