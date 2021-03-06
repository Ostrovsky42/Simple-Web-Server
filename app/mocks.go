// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package app is a generated GoMock package.
package app

import (
	library "SimplyWebServer/library"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AddTransaction mocks base method.
func (m *MockStorage) AddTransaction(transactType library.TransactionType, userId int, amount float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTransaction", transactType, userId, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTransaction indicates an expected call of AddTransaction.
func (mr *MockStorageMockRecorder) AddTransaction(transactType, userId, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransaction", reflect.TypeOf((*MockStorage)(nil).AddTransaction), transactType, userId, amount)
}

// CreateUser mocks base method.
func (m *MockStorage) CreateUser() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStorageMockRecorder) CreateUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStorage)(nil).CreateUser))
}

// UpdateBalanceByUserId mocks base method.
func (m *MockStorage) UpdateBalanceByUserId(userId int, amount float64) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBalanceByUserId", userId, amount)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateBalanceByUserId indicates an expected call of UpdateBalanceByUserId.
func (mr *MockStorageMockRecorder) UpdateBalanceByUserId(userId, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBalanceByUserId", reflect.TypeOf((*MockStorage)(nil).UpdateBalanceByUserId), userId, amount)
}
