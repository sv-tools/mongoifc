// Code generated by MockGen. DO NOT EDIT.
// Source: ./change_stream.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
)

// MockChangeStream is a mock of ChangeStream interface.
type MockChangeStream struct {
	ctrl     *gomock.Controller
	recorder *MockChangeStreamMockRecorder
}

// MockChangeStreamMockRecorder is the mock recorder for MockChangeStream.
type MockChangeStreamMockRecorder struct {
	mock *MockChangeStream
}

// NewMockChangeStream creates a new mock instance.
func NewMockChangeStream(ctrl *gomock.Controller) *MockChangeStream {
	mock := &MockChangeStream{ctrl: ctrl}
	mock.recorder = &MockChangeStreamMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChangeStream) EXPECT() *MockChangeStreamMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockChangeStream) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockChangeStreamMockRecorder) Close(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockChangeStream)(nil).Close), ctx)
}

// Current mocks base method.
func (m *MockChangeStream) Current() bson.Raw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Current")
	ret0, _ := ret[0].(bson.Raw)
	return ret0
}

// Current indicates an expected call of Current.
func (mr *MockChangeStreamMockRecorder) Current() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Current", reflect.TypeOf((*MockChangeStream)(nil).Current))
}

// Decode mocks base method.
func (m *MockChangeStream) Decode(val interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", val)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode.
func (mr *MockChangeStreamMockRecorder) Decode(val interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockChangeStream)(nil).Decode), val)
}

// Err mocks base method.
func (m *MockChangeStream) Err() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err.
func (mr *MockChangeStreamMockRecorder) Err() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Err", reflect.TypeOf((*MockChangeStream)(nil).Err))
}

// ID mocks base method.
func (m *MockChangeStream) ID() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(int64)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockChangeStreamMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockChangeStream)(nil).ID))
}

// Next mocks base method.
func (m *MockChangeStream) Next(ctx context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next", ctx)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockChangeStreamMockRecorder) Next(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockChangeStream)(nil).Next), ctx)
}

// ResumeToken mocks base method.
func (m *MockChangeStream) ResumeToken() bson.Raw {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResumeToken")
	ret0, _ := ret[0].(bson.Raw)
	return ret0
}

// ResumeToken indicates an expected call of ResumeToken.
func (mr *MockChangeStreamMockRecorder) ResumeToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResumeToken", reflect.TypeOf((*MockChangeStream)(nil).ResumeToken))
}

// TryNext mocks base method.
func (m *MockChangeStream) TryNext(ctx context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TryNext", ctx)
	ret0, _ := ret[0].(bool)
	return ret0
}

// TryNext indicates an expected call of TryNext.
func (mr *MockChangeStreamMockRecorder) TryNext(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TryNext", reflect.TypeOf((*MockChangeStream)(nil).TryNext), ctx)
}

// WrappedChangeStream mocks base method.
func (m *MockChangeStream) WrappedChangeStream() *mongo.ChangeStream {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WrappedChangeStream")
	ret0, _ := ret[0].(*mongo.ChangeStream)
	return ret0
}

// WrappedChangeStream indicates an expected call of WrappedChangeStream.
func (mr *MockChangeStreamMockRecorder) WrappedChangeStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WrappedChangeStream", reflect.TypeOf((*MockChangeStream)(nil).WrappedChangeStream))
}