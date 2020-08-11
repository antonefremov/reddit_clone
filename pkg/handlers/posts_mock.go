// Code generated by MockGen. DO NOT EDIT.
// Source: posts.go

// Package handlers is a generated GoMock package.
package handlers

import (
	gomock "github.com/golang/mock/gomock"
	posts "golang-stepik-2020q2/6/99_hw/redditclone/pkg/posts"
	reflect "reflect"
)

// MockPostsRepoInterface is a mock of PostsRepoInterface interface
type MockPostsRepoInterface struct {
	ctrl     *gomock.Controller
	recorder *MockPostsRepoInterfaceMockRecorder
}

// MockPostsRepoInterfaceMockRecorder is the mock recorder for MockPostsRepoInterface
type MockPostsRepoInterfaceMockRecorder struct {
	mock *MockPostsRepoInterface
}

// NewMockPostsRepoInterface creates a new mock instance
func NewMockPostsRepoInterface(ctrl *gomock.Controller) *MockPostsRepoInterface {
	mock := &MockPostsRepoInterface{ctrl: ctrl}
	mock.recorder = &MockPostsRepoInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPostsRepoInterface) EXPECT() *MockPostsRepoInterfaceMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockPostsRepoInterface) All() ([]*posts.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].([]*posts.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// All indicates an expected call of All
func (mr *MockPostsRepoInterfaceMockRecorder) All() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockPostsRepoInterface)(nil).All))
}

// ListByCategory mocks base method
func (m *MockPostsRepoInterface) ListByCategory(arg0 string) ([]*posts.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByCategory", arg0)
	ret0, _ := ret[0].([]*posts.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByCategory indicates an expected call of ListByCategory
func (mr *MockPostsRepoInterfaceMockRecorder) ListByCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByCategory", reflect.TypeOf((*MockPostsRepoInterface)(nil).ListByCategory), arg0)
}

// Get mocks base method
func (m *MockPostsRepoInterface) Get(arg0 string) (*posts.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*posts.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockPostsRepoInterfaceMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPostsRepoInterface)(nil).Get), arg0)
}

// GetByAuthor mocks base method
func (m *MockPostsRepoInterface) GetByAuthor(arg0 string) ([]*posts.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAuthor", arg0)
	ret0, _ := ret[0].([]*posts.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByAuthor indicates an expected call of GetByAuthor
func (mr *MockPostsRepoInterfaceMockRecorder) GetByAuthor(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAuthor", reflect.TypeOf((*MockPostsRepoInterface)(nil).GetByAuthor), arg0)
}

// Add mocks base method
func (m *MockPostsRepoInterface) Add(arg0 *posts.Post) (*posts.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(*posts.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockPostsRepoInterfaceMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPostsRepoInterface)(nil).Add), arg0)
}

// AddComment mocks base method
func (m *MockPostsRepoInterface) AddComment(arg0 string, arg1 *posts.Comment) (*posts.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddComment", arg0, arg1)
	ret0, _ := ret[0].(*posts.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddComment indicates an expected call of AddComment
func (mr *MockPostsRepoInterfaceMockRecorder) AddComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddComment", reflect.TypeOf((*MockPostsRepoInterface)(nil).AddComment), arg0, arg1)
}

// DeleteComment mocks base method
func (m *MockPostsRepoInterface) DeleteComment(arg0, arg1 string) (*posts.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", arg0, arg1)
	ret0, _ := ret[0].(*posts.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteComment indicates an expected call of DeleteComment
func (mr *MockPostsRepoInterfaceMockRecorder) DeleteComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockPostsRepoInterface)(nil).DeleteComment), arg0, arg1)
}

// Delete mocks base method
func (m *MockPostsRepoInterface) Delete(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockPostsRepoInterfaceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPostsRepoInterface)(nil).Delete), arg0, arg1)
}

// Vote mocks base method
func (m *MockPostsRepoInterface) Vote(arg0 string, arg1 posts.Vote) (*posts.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vote", arg0, arg1)
	ret0, _ := ret[0].(*posts.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Vote indicates an expected call of Vote
func (mr *MockPostsRepoInterfaceMockRecorder) Vote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vote", reflect.TypeOf((*MockPostsRepoInterface)(nil).Vote), arg0, arg1)
}

// Unvote mocks base method
func (m *MockPostsRepoInterface) Unvote(arg0, arg1 string) (*posts.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unvote", arg0, arg1)
	ret0, _ := ret[0].(*posts.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Unvote indicates an expected call of Unvote
func (mr *MockPostsRepoInterfaceMockRecorder) Unvote(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unvote", reflect.TypeOf((*MockPostsRepoInterface)(nil).Unvote), arg0, arg1)
}