// Code generated by MockGen. DO NOT EDIT.
// Source: employers/internal/service/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	autogen "github.com/qaZar1/HHforURFU/employers/autogen"
)

// MockServiceInterface is a mock of ServiceInterface interface.
type MockServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServiceInterfaceMockRecorder
}

// MockServiceInterfaceMockRecorder is the mock recorder for MockServiceInterface.
type MockServiceInterfaceMockRecorder struct {
	mock *MockServiceInterface
}

// NewMockServiceInterface creates a new mock instance.
func NewMockServiceInterface(ctrl *gomock.Controller) *MockServiceInterface {
	mock := &MockServiceInterface{ctrl: ctrl}
	mock.recorder = &MockServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceInterface) EXPECT() *MockServiceInterfaceMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockServiceInterface) AddUser(user autogen.User) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", user)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddUser indicates an expected call of AddUser.
func (mr *MockServiceInterfaceMockRecorder) AddUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockServiceInterface)(nil).AddUser), user)
}

// GetAllUsers mocks base method.
func (m *MockServiceInterface) GetAllUsers() ([]autogen.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]autogen.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockServiceInterfaceMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockServiceInterface)(nil).GetAllUsers))
}

// GetUserByChatID mocks base method.
func (m *MockServiceInterface) GetUserByChatID(chatId int64) (*autogen.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByChatID", chatId)
	ret0, _ := ret[0].(*autogen.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByChatID indicates an expected call of GetUserByChatID.
func (mr *MockServiceInterfaceMockRecorder) GetUserByChatID(chatId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByChatID", reflect.TypeOf((*MockServiceInterface)(nil).GetUserByChatID), chatId)
}

// RemoveUser mocks base method.
func (m *MockServiceInterface) RemoveUser(chatId int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", chatId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockServiceInterfaceMockRecorder) RemoveUser(chatId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockServiceInterface)(nil).RemoveUser), chatId)
}

// UpdateUser mocks base method.
func (m *MockServiceInterface) UpdateUser(chatId int64, updateUser autogen.UpdateUser) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", chatId, updateUser)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockServiceInterfaceMockRecorder) UpdateUser(chatId, updateUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockServiceInterface)(nil).UpdateUser), chatId, updateUser)
}