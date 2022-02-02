// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/common (interfaces: LeadershipPinningBackend,LeadershipMachine)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/juju/juju/apiserver/common"
)

// MockLeadershipPinningBackend is a mock of LeadershipPinningBackend interface.
type MockLeadershipPinningBackend struct {
	ctrl     *gomock.Controller
	recorder *MockLeadershipPinningBackendMockRecorder
}

// MockLeadershipPinningBackendMockRecorder is the mock recorder for MockLeadershipPinningBackend.
type MockLeadershipPinningBackendMockRecorder struct {
	mock *MockLeadershipPinningBackend
}

// NewMockLeadershipPinningBackend creates a new mock instance.
func NewMockLeadershipPinningBackend(ctrl *gomock.Controller) *MockLeadershipPinningBackend {
	mock := &MockLeadershipPinningBackend{ctrl: ctrl}
	mock.recorder = &MockLeadershipPinningBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeadershipPinningBackend) EXPECT() *MockLeadershipPinningBackendMockRecorder {
	return m.recorder
}

// Machine mocks base method.
func (m *MockLeadershipPinningBackend) Machine(arg0 string) (common.LeadershipMachine, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(common.LeadershipMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine.
func (mr *MockLeadershipPinningBackendMockRecorder) Machine(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockLeadershipPinningBackend)(nil).Machine), arg0)
}

// MockLeadershipMachine is a mock of LeadershipMachine interface.
type MockLeadershipMachine struct {
	ctrl     *gomock.Controller
	recorder *MockLeadershipMachineMockRecorder
}

// MockLeadershipMachineMockRecorder is the mock recorder for MockLeadershipMachine.
type MockLeadershipMachineMockRecorder struct {
	mock *MockLeadershipMachine
}

// NewMockLeadershipMachine creates a new mock instance.
func NewMockLeadershipMachine(ctrl *gomock.Controller) *MockLeadershipMachine {
	mock := &MockLeadershipMachine{ctrl: ctrl}
	mock.recorder = &MockLeadershipMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLeadershipMachine) EXPECT() *MockLeadershipMachineMockRecorder {
	return m.recorder
}

// ApplicationNames mocks base method.
func (m *MockLeadershipMachine) ApplicationNames() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationNames")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ApplicationNames indicates an expected call of ApplicationNames.
func (mr *MockLeadershipMachineMockRecorder) ApplicationNames() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationNames", reflect.TypeOf((*MockLeadershipMachine)(nil).ApplicationNames))
}
