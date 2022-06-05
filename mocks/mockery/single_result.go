// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	bson "go.mongodb.org/mongo-driver/bson"
)

// SingleResult is an autogenerated mock type for the SingleResult type
type SingleResult struct {
	mock.Mock
}

// Decode provides a mock function with given fields: v
func (_m *SingleResult) Decode(v interface{}) error {
	ret := _m.Called(v)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DecodeBytes provides a mock function with given fields:
func (_m *SingleResult) DecodeBytes() (bson.Raw, error) {
	ret := _m.Called()

	var r0 bson.Raw
	if rf, ok := ret.Get(0).(func() bson.Raw); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(bson.Raw)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Err provides a mock function with given fields:
func (_m *SingleResult) Err() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewSingleResultT interface {
	mock.TestingT
	Cleanup(func())
}

// NewSingleResult creates a new instance of SingleResult. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSingleResult(t NewSingleResultT) *SingleResult {
	mock := &SingleResult{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
