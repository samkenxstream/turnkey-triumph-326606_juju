// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facade (interfaces: Authorizer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	permission "github.com/juju/juju/core/permission"
	names "github.com/juju/names/v4"
)

// MockAuthorizer is a mock of Authorizer interface.
type MockAuthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizerMockRecorder
}

// MockAuthorizerMockRecorder is the mock recorder for MockAuthorizer.
type MockAuthorizerMockRecorder struct {
	mock *MockAuthorizer
}

// NewMockAuthorizer creates a new mock instance.
func NewMockAuthorizer(ctrl *gomock.Controller) *MockAuthorizer {
	mock := &MockAuthorizer{ctrl: ctrl}
	mock.recorder = &MockAuthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizer) EXPECT() *MockAuthorizerMockRecorder {
	return m.recorder
}

// AuthApplicationAgent mocks base method.
func (m *MockAuthorizer) AuthApplicationAgent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthApplicationAgent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthApplicationAgent indicates an expected call of AuthApplicationAgent.
func (mr *MockAuthorizerMockRecorder) AuthApplicationAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthApplicationAgent", reflect.TypeOf((*MockAuthorizer)(nil).AuthApplicationAgent))
}

// AuthClient mocks base method.
func (m *MockAuthorizer) AuthClient() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthClient")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthClient indicates an expected call of AuthClient.
func (mr *MockAuthorizerMockRecorder) AuthClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthClient", reflect.TypeOf((*MockAuthorizer)(nil).AuthClient))
}

// AuthController mocks base method.
func (m *MockAuthorizer) AuthController() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthController")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthController indicates an expected call of AuthController.
func (mr *MockAuthorizerMockRecorder) AuthController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthController", reflect.TypeOf((*MockAuthorizer)(nil).AuthController))
}

// AuthMachineAgent mocks base method.
func (m *MockAuthorizer) AuthMachineAgent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthMachineAgent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthMachineAgent indicates an expected call of AuthMachineAgent.
func (mr *MockAuthorizerMockRecorder) AuthMachineAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthMachineAgent", reflect.TypeOf((*MockAuthorizer)(nil).AuthMachineAgent))
}

// AuthModelAgent mocks base method.
func (m *MockAuthorizer) AuthModelAgent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthModelAgent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthModelAgent indicates an expected call of AuthModelAgent.
func (mr *MockAuthorizerMockRecorder) AuthModelAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthModelAgent", reflect.TypeOf((*MockAuthorizer)(nil).AuthModelAgent))
}

// AuthOwner mocks base method.
func (m *MockAuthorizer) AuthOwner(arg0 names.Tag) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthOwner", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthOwner indicates an expected call of AuthOwner.
func (mr *MockAuthorizerMockRecorder) AuthOwner(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthOwner", reflect.TypeOf((*MockAuthorizer)(nil).AuthOwner), arg0)
}

// AuthTokenString mocks base method.
func (m *MockAuthorizer) AuthTokenString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthTokenString")
	ret0, _ := ret[0].(string)
	return ret0
}

// AuthTokenString indicates an expected call of AuthTokenString.
func (mr *MockAuthorizerMockRecorder) AuthTokenString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthTokenString", reflect.TypeOf((*MockAuthorizer)(nil).AuthTokenString))
}

// AuthUnitAgent mocks base method.
func (m *MockAuthorizer) AuthUnitAgent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthUnitAgent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthUnitAgent indicates an expected call of AuthUnitAgent.
func (mr *MockAuthorizerMockRecorder) AuthUnitAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthUnitAgent", reflect.TypeOf((*MockAuthorizer)(nil).AuthUnitAgent))
}

// ConnectedModel mocks base method.
func (m *MockAuthorizer) ConnectedModel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectedModel")
	ret0, _ := ret[0].(string)
	return ret0
}

// ConnectedModel indicates an expected call of ConnectedModel.
func (mr *MockAuthorizerMockRecorder) ConnectedModel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectedModel", reflect.TypeOf((*MockAuthorizer)(nil).ConnectedModel))
}

// EntityHasPermission mocks base method.
func (m *MockAuthorizer) EntityHasPermission(arg0 names.Tag, arg1 permission.Access, arg2 names.Tag) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EntityHasPermission", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EntityHasPermission indicates an expected call of EntityHasPermission.
func (mr *MockAuthorizerMockRecorder) EntityHasPermission(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntityHasPermission", reflect.TypeOf((*MockAuthorizer)(nil).EntityHasPermission), arg0, arg1, arg2)
}

// GetAuthTag mocks base method.
func (m *MockAuthorizer) GetAuthTag() names.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthTag")
	ret0, _ := ret[0].(names.Tag)
	return ret0
}

// GetAuthTag indicates an expected call of GetAuthTag.
func (mr *MockAuthorizerMockRecorder) GetAuthTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthTag", reflect.TypeOf((*MockAuthorizer)(nil).GetAuthTag))
}

// HasPermission mocks base method.
func (m *MockAuthorizer) HasPermission(arg0 permission.Access, arg1 names.Tag) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPermission", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasPermission indicates an expected call of HasPermission.
func (mr *MockAuthorizerMockRecorder) HasPermission(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPermission", reflect.TypeOf((*MockAuthorizer)(nil).HasPermission), arg0, arg1)
}
