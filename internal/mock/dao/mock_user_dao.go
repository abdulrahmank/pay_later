// Code generated by MockGen. DO NOT EDIT.
// Source: ./dao/user_dao.go

// Package dao is a generated GoMock package.
package dao

import (
	model "github.com/abdulrahmank/pay_later/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserDao is a mock of UserDao interface
type MockUserDao struct {
	ctrl     *gomock.Controller
	recorder *MockUserDaoMockRecorder
}

// MockUserDaoMockRecorder is the mock recorder for MockUserDao
type MockUserDaoMockRecorder struct {
	mock *MockUserDao
}

// NewMockUserDao creates a new mock instance
func NewMockUserDao(ctrl *gomock.Controller) *MockUserDao {
	mock := &MockUserDao{ctrl: ctrl}
	mock.recorder = &MockUserDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserDao) EXPECT() *MockUserDaoMockRecorder {
	return m.recorder
}

// SaveUser mocks base method
func (m *MockUserDao) SaveUser(name, mailId string, creditLimit float32) error {
	ret := m.ctrl.Call(m, "SaveUser", name, mailId, creditLimit)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveUser indicates an expected call of SaveUser
func (mr *MockUserDaoMockRecorder) SaveUser(name, mailId, creditLimit interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveUser", reflect.TypeOf((*MockUserDao)(nil).SaveUser), name, mailId, creditLimit)
}

// GetUser mocks base method
func (m *MockUserDao) GetUser(name string) *model.User {
	ret := m.ctrl.Call(m, "GetUser", name)
	ret0, _ := ret[0].(*model.User)
	return ret0
}

// GetUser indicates an expected call of GetUser
func (mr *MockUserDaoMockRecorder) GetUser(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserDao)(nil).GetUser), name)
}

// IncrementDues mocks base method
func (m *MockUserDao) IncrementDues(user *model.User, txnAmt float32) error {
	ret := m.ctrl.Call(m, "IncrementDues", user, txnAmt)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementDues indicates an expected call of IncrementDues
func (mr *MockUserDaoMockRecorder) IncrementDues(user, txnAmt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementDues", reflect.TypeOf((*MockUserDao)(nil).IncrementDues), user, txnAmt)
}

// GetAllUsers mocks base method
func (m *MockUserDao) GetAllUsers() []*model.User {
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]*model.User)
	return ret0
}

// GetAllUsers indicates an expected call of GetAllUsers
func (mr *MockUserDaoMockRecorder) GetAllUsers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserDao)(nil).GetAllUsers))
}
