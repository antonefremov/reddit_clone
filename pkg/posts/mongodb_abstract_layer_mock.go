// Code generated by MockGen. DO NOT EDIT.
// Source: mongodb_abstract_layer.go

// Package posts is a generated GoMock package.
package posts

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockIMongoDatabase is a mock of IMongoDatabase interface
type MockIMongoDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockIMongoDatabaseMockRecorder
}

// MockIMongoDatabaseMockRecorder is the mock recorder for MockIMongoDatabase
type MockIMongoDatabaseMockRecorder struct {
	mock *MockIMongoDatabase
}

// NewMockIMongoDatabase creates a new mock instance
func NewMockIMongoDatabase(ctrl *gomock.Controller) *MockIMongoDatabase {
	mock := &MockIMongoDatabase{ctrl: ctrl}
	mock.recorder = &MockIMongoDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIMongoDatabase) EXPECT() *MockIMongoDatabaseMockRecorder {
	return m.recorder
}

// Collection mocks base method
func (m *MockIMongoDatabase) Collection(name string) IMongoCollection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Collection", name)
	ret0, _ := ret[0].(IMongoCollection)
	return ret0
}

// Collection indicates an expected call of Collection
func (mr *MockIMongoDatabaseMockRecorder) Collection(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Collection", reflect.TypeOf((*MockIMongoDatabase)(nil).Collection), name)
}

// MockIMongoCollection is a mock of IMongoCollection interface
type MockIMongoCollection struct {
	ctrl     *gomock.Controller
	recorder *MockIMongoCollectionMockRecorder
}

// MockIMongoCollectionMockRecorder is the mock recorder for MockIMongoCollection
type MockIMongoCollectionMockRecorder struct {
	mock *MockIMongoCollection
}

// NewMockIMongoCollection creates a new mock instance
func NewMockIMongoCollection(ctrl *gomock.Controller) *MockIMongoCollection {
	mock := &MockIMongoCollection{ctrl: ctrl}
	mock.recorder = &MockIMongoCollectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIMongoCollection) EXPECT() *MockIMongoCollectionMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockIMongoCollection) Find(ctx context.Context, filter interface{}) (IMongoCursor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", ctx, filter)
	ret0, _ := ret[0].(IMongoCursor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find
func (mr *MockIMongoCollectionMockRecorder) Find(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockIMongoCollection)(nil).Find), ctx, filter)
}

// FindOne mocks base method
func (m *MockIMongoCollection) FindOne(ctx context.Context, filter interface{}) IMongoSingleResult {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", ctx, filter)
	ret0, _ := ret[0].(IMongoSingleResult)
	return ret0
}

// FindOne indicates an expected call of FindOne
func (mr *MockIMongoCollectionMockRecorder) FindOne(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockIMongoCollection)(nil).FindOne), ctx, filter)
}

// InsertOne mocks base method
func (m *MockIMongoCollection) InsertOne(ctx context.Context, item interface{}) (IMongoInsertOneResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOne", ctx, item)
	ret0, _ := ret[0].(IMongoInsertOneResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertOne indicates an expected call of InsertOne
func (mr *MockIMongoCollectionMockRecorder) InsertOne(ctx, item interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockIMongoCollection)(nil).InsertOne), ctx, item)
}

// DeleteOne mocks base method
func (m *MockIMongoCollection) DeleteOne(ctx context.Context, filter interface{}) (IMongoDeleteResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOne", ctx, filter)
	ret0, _ := ret[0].(IMongoDeleteResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteOne indicates an expected call of DeleteOne
func (mr *MockIMongoCollectionMockRecorder) DeleteOne(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOne", reflect.TypeOf((*MockIMongoCollection)(nil).DeleteOne), ctx, filter)
}

// ReplaceOne mocks base method
func (m *MockIMongoCollection) ReplaceOne(ctx context.Context, filter, replacement interface{}) (IMongoUpdateResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReplaceOne", ctx, filter, replacement)
	ret0, _ := ret[0].(IMongoUpdateResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReplaceOne indicates an expected call of ReplaceOne
func (mr *MockIMongoCollectionMockRecorder) ReplaceOne(ctx, filter, replacement interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplaceOne", reflect.TypeOf((*MockIMongoCollection)(nil).ReplaceOne), ctx, filter, replacement)
}

// MockIMongoSingleResult is a mock of IMongoSingleResult interface
type MockIMongoSingleResult struct {
	ctrl     *gomock.Controller
	recorder *MockIMongoSingleResultMockRecorder
}

// MockIMongoSingleResultMockRecorder is the mock recorder for MockIMongoSingleResult
type MockIMongoSingleResultMockRecorder struct {
	mock *MockIMongoSingleResult
}

// NewMockIMongoSingleResult creates a new mock instance
func NewMockIMongoSingleResult(ctrl *gomock.Controller) *MockIMongoSingleResult {
	mock := &MockIMongoSingleResult{ctrl: ctrl}
	mock.recorder = &MockIMongoSingleResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIMongoSingleResult) EXPECT() *MockIMongoSingleResultMockRecorder {
	return m.recorder
}

// Decode mocks base method
func (m *MockIMongoSingleResult) Decode(v interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", v)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode
func (mr *MockIMongoSingleResultMockRecorder) Decode(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockIMongoSingleResult)(nil).Decode), v)
}

// MockIMongoCursor is a mock of IMongoCursor interface
type MockIMongoCursor struct {
	ctrl     *gomock.Controller
	recorder *MockIMongoCursorMockRecorder
}

// MockIMongoCursorMockRecorder is the mock recorder for MockIMongoCursor
type MockIMongoCursorMockRecorder struct {
	mock *MockIMongoCursor
}

// NewMockIMongoCursor creates a new mock instance
func NewMockIMongoCursor(ctrl *gomock.Controller) *MockIMongoCursor {
	mock := &MockIMongoCursor{ctrl: ctrl}
	mock.recorder = &MockIMongoCursorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIMongoCursor) EXPECT() *MockIMongoCursorMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockIMongoCursor) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockIMongoCursorMockRecorder) Close(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIMongoCursor)(nil).Close), arg0)
}

// Next mocks base method
func (m *MockIMongoCursor) Next(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockIMongoCursorMockRecorder) Next(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockIMongoCursor)(nil).Next), arg0)
}

// Decode mocks base method
func (m *MockIMongoCursor) Decode(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode
func (mr *MockIMongoCursorMockRecorder) Decode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockIMongoCursor)(nil).Decode), arg0)
}

// MockIMongoInsertOneResult is a mock of IMongoInsertOneResult interface
type MockIMongoInsertOneResult struct {
	ctrl     *gomock.Controller
	recorder *MockIMongoInsertOneResultMockRecorder
}

// MockIMongoInsertOneResultMockRecorder is the mock recorder for MockIMongoInsertOneResult
type MockIMongoInsertOneResultMockRecorder struct {
	mock *MockIMongoInsertOneResult
}

// NewMockIMongoInsertOneResult creates a new mock instance
func NewMockIMongoInsertOneResult(ctrl *gomock.Controller) *MockIMongoInsertOneResult {
	mock := &MockIMongoInsertOneResult{ctrl: ctrl}
	mock.recorder = &MockIMongoInsertOneResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIMongoInsertOneResult) EXPECT() *MockIMongoInsertOneResultMockRecorder {
	return m.recorder
}

// MockIMongoDeleteResult is a mock of IMongoDeleteResult interface
type MockIMongoDeleteResult struct {
	ctrl     *gomock.Controller
	recorder *MockIMongoDeleteResultMockRecorder
}

// MockIMongoDeleteResultMockRecorder is the mock recorder for MockIMongoDeleteResult
type MockIMongoDeleteResultMockRecorder struct {
	mock *MockIMongoDeleteResult
}

// NewMockIMongoDeleteResult creates a new mock instance
func NewMockIMongoDeleteResult(ctrl *gomock.Controller) *MockIMongoDeleteResult {
	mock := &MockIMongoDeleteResult{ctrl: ctrl}
	mock.recorder = &MockIMongoDeleteResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIMongoDeleteResult) EXPECT() *MockIMongoDeleteResultMockRecorder {
	return m.recorder
}

// MockIMongoUpdateResult is a mock of IMongoUpdateResult interface
type MockIMongoUpdateResult struct {
	ctrl     *gomock.Controller
	recorder *MockIMongoUpdateResultMockRecorder
}

// MockIMongoUpdateResultMockRecorder is the mock recorder for MockIMongoUpdateResult
type MockIMongoUpdateResultMockRecorder struct {
	mock *MockIMongoUpdateResult
}

// NewMockIMongoUpdateResult creates a new mock instance
func NewMockIMongoUpdateResult(ctrl *gomock.Controller) *MockIMongoUpdateResult {
	mock := &MockIMongoUpdateResult{ctrl: ctrl}
	mock.recorder = &MockIMongoUpdateResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIMongoUpdateResult) EXPECT() *MockIMongoUpdateResultMockRecorder {
	return m.recorder
}
