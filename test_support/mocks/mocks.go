// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudfoundry/bosh-cli/test_support (interfaces: Spy)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSpy is a mock of Spy interface.
type MockSpy struct {
	ctrl     *gomock.Controller
	recorder *MockSpyMockRecorder
}

// MockSpyMockRecorder is the mock recorder for MockSpy.
type MockSpyMockRecorder struct {
	mock *MockSpy
}

// NewMockSpy creates a new mock instance.
func NewMockSpy(ctrl *gomock.Controller) *MockSpy {
	mock := &MockSpy{ctrl: ctrl}
	mock.recorder = &MockSpyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpy) EXPECT() *MockSpyMockRecorder {
	return m.recorder
}

// Record mocks base method.
func (m *MockSpy) Record() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Record")
}

// Record indicates an expected call of Record.
func (mr *MockSpyMockRecorder) Record() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Record", reflect.TypeOf((*MockSpy)(nil).Record))
}
