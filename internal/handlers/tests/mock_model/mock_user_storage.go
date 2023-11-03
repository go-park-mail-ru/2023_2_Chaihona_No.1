// Code generated by MockGen. DO NOT EDIT.
// Source: ../../repositories/users/user_repository.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	reflect "reflect"

	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// ChangeUser mocks base method.
func (m *MockUserRepository) ChangeUser(user model.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeUser indicates an expected call of ChangeUser.
func (mr *MockUserRepositoryMockRecorder) ChangeUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUser", reflect.TypeOf((*MockUserRepository)(nil).ChangeUser), user)
}

// CheckUser mocks base method.
func (m *MockUserRepository) CheckUser(login string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUser", login)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUser indicates an expected call of CheckUser.
func (mr *MockUserRepositoryMockRecorder) CheckUser(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUser", reflect.TypeOf((*MockUserRepository)(nil).CheckUser), login)
}

// DeleteUser mocks base method.
func (m *MockUserRepository) DeleteUser(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryMockRecorder) DeleteUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser), id)
}

// GetUser mocks base method.
func (m *MockUserRepository) GetUser(id int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserRepositoryMockRecorder) GetUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepository)(nil).GetUser), id)
}

// GetUserWithSubscribers mocks base method.
func (m *MockUserRepository) GetUserWithSubscribers(id int) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserWithSubscribers", id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserWithSubscribers indicates an expected call of GetUserWithSubscribers.
func (mr *MockUserRepositoryMockRecorder) GetUserWithSubscribers(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserWithSubscribers", reflect.TypeOf((*MockUserRepository)(nil).GetUserWithSubscribers), id)
}

// GetUsers mocks base method.
func (m *MockUserRepository) GetUsers() ([]model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers.
func (mr *MockUserRepositoryMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserRepository)(nil).GetUsers))
}

// RegisterNewUser mocks base method.
func (m *MockUserRepository) RegisterNewUser(user *model.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterNewUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterNewUser indicates an expected call of RegisterNewUser.
func (mr *MockUserRepositoryMockRecorder) RegisterNewUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterNewUser", reflect.TypeOf((*MockUserRepository)(nil).RegisterNewUser), user)
}