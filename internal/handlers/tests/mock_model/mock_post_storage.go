// Code generated by MockGen. DO NOT EDIT.
// Source: ../../repositories/posts/post_repository.go

// Package mock_model is a generated GoMock package.
package mock_model

import (
	reflect "reflect"

	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockPostRepository is a mock of PostRepository interface.
type MockPostRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPostRepositoryMockRecorder
}

// MockPostRepositoryMockRecorder is the mock recorder for MockPostRepository.
type MockPostRepositoryMockRecorder struct {
	mock *MockPostRepository
}

// NewMockPostRepository creates a new mock instance.
func NewMockPostRepository(ctrl *gomock.Controller) *MockPostRepository {
	mock := &MockPostRepository{ctrl: ctrl}
	mock.recorder = &MockPostRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPostRepository) EXPECT() *MockPostRepositoryMockRecorder {
	return m.recorder
}

// ChangePost mocks base method.
func (m *MockPostRepository) ChangePost(post model.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePost", post)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePost indicates an expected call of ChangePost.
func (mr *MockPostRepositoryMockRecorder) ChangePost(post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePost", reflect.TypeOf((*MockPostRepository)(nil).ChangePost), post)
}

// CreateNewPost mocks base method.
func (m *MockPostRepository) CreateNewPost(post model.Post) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewPost", post)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewPost indicates an expected call of CreateNewPost.
func (mr *MockPostRepositoryMockRecorder) CreateNewPost(post interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewPost", reflect.TypeOf((*MockPostRepository)(nil).CreateNewPost), post)
}

// DeletePost mocks base method.
func (m *MockPostRepository) DeletePost(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockPostRepositoryMockRecorder) DeletePost(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockPostRepository)(nil).DeletePost), id)
}

// GetPostById mocks base method.
func (m *MockPostRepository) GetPostById(id uint) (model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostById", id)
	ret0, _ := ret[0].(model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostById indicates an expected call of GetPostById.
func (mr *MockPostRepositoryMockRecorder) GetPostById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostById", reflect.TypeOf((*MockPostRepository)(nil).GetPostById), id)
}

// GetPostsByAuthorId mocks base method.
func (m *MockPostRepository) GetPostsByAuthorId(authorID, subscriberId uint) ([]model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsByAuthorId", authorID, subscriberId)
	ret0, _ := ret[0].([]model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsByAuthorId indicates an expected call of GetPostsByAuthorId.
func (mr *MockPostRepositoryMockRecorder) GetPostsByAuthorId(authorID, subscriberId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsByAuthorId", reflect.TypeOf((*MockPostRepository)(nil).GetPostsByAuthorId), authorID, subscriberId)
}

// GetUsersFeed mocks base method.
func (m *MockPostRepository) GetUsersFeed(userId uint) ([]model.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsersFeed", userId)
	ret0, _ := ret[0].([]model.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsersFeed indicates an expected call of GetUsersFeed.
func (mr *MockPostRepositoryMockRecorder) GetUsersFeed(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsersFeed", reflect.TypeOf((*MockPostRepository)(nil).GetUsersFeed), userId)
}
