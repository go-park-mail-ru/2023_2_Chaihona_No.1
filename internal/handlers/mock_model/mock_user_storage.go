// Code generated by MockGen. DO NOT EDIT.
// Source: ../../registration/user_repository.go

// Package mock_registration is a generated GoMock package.
package mock_model

import (
	reflect "reflect"

	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
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

// CheckUser mocks base method.
func (m *MockUserRepository) CheckUser(login string) (*model.User, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUser", login)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// CheckUser indicates an expected call of CheckUser.
func (mr *MockUserRepositoryMockRecorder) CheckUser(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUser", reflect.TypeOf((*MockUserRepository)(nil).CheckUser), login)
}

// DeleteUser mocks base method.
func (m *MockUserRepository) DeleteUser(login string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", login)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserRepositoryMockRecorder) DeleteUser(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserRepository)(nil).DeleteUser), login)
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
func (m *MockUserRepository) RegisterNewUser(user *model.User) *users.ErrorUserRegistration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterNewUser", user)
	ret0, _ := ret[0].(*users.ErrorUserRegistration)
	return ret0
}

// RegisterNewUser indicates an expected call of RegisterNewUser.
func (mr *MockUserRepositoryMockRecorder) RegisterNewUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterNewUser", reflect.TypeOf((*MockUserRepository)(nil).RegisterNewUser), user)
}