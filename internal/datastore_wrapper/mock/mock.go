// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/datastore_wrapper/datastore_wrapper.go

// Package mock_datastore_wrapper is a generated GoMock package.
package mock_datastore_wrapper

import (
	datastore "cloud.google.com/go/datastore"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDatastoreWrapper is a mock of DatastoreWrapper interface
type MockDatastoreWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockDatastoreWrapperMockRecorder
}

// MockDatastoreWrapperMockRecorder is the mock recorder for MockDatastoreWrapper
type MockDatastoreWrapperMockRecorder struct {
	mock *MockDatastoreWrapper
}

// NewMockDatastoreWrapper creates a new mock instance
func NewMockDatastoreWrapper(ctrl *gomock.Controller) *MockDatastoreWrapper {
	mock := &MockDatastoreWrapper{ctrl: ctrl}
	mock.recorder = &MockDatastoreWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatastoreWrapper) EXPECT() *MockDatastoreWrapperMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockDatastoreWrapper) Get(arg0 context.Context, arg1 *datastore.Key, arg2 interface{}) error {
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockDatastoreWrapperMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDatastoreWrapper)(nil).Get), arg0, arg1, arg2)
}

// GetAll mocks base method
func (m *MockDatastoreWrapper) GetAll(arg0 context.Context, arg1 *datastore.Query, arg2 interface{}) ([]*datastore.Key, error) {
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*datastore.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockDatastoreWrapperMockRecorder) GetAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockDatastoreWrapper)(nil).GetAll), arg0, arg1, arg2)
}

// Put mocks base method
func (m *MockDatastoreWrapper) Put(arg0 context.Context, arg1 *datastore.Key, arg2 interface{}) (*datastore.Key, error) {
	ret := m.ctrl.Call(m, "Put", arg0, arg1, arg2)
	ret0, _ := ret[0].(*datastore.Key)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Put indicates an expected call of Put
func (mr *MockDatastoreWrapperMockRecorder) Put(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockDatastoreWrapper)(nil).Put), arg0, arg1, arg2)
}
