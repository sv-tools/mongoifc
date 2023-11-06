// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	mongo "go.mongodb.org/mongo-driver/mongo"

	mongoifc "github.com/sv-tools/mongoifc"

	options "go.mongodb.org/mongo-driver/mongo/options"

	readpref "go.mongodb.org/mongo-driver/mongo/readpref"

	time "time"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

type Client_Expecter struct {
	mock *mock.Mock
}

func (_m *Client) EXPECT() *Client_Expecter {
	return &Client_Expecter{mock: &_m.Mock}
}

// Connect provides a mock function with given fields: ctx
func (_m *Client) Connect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type Client_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Client_Expecter) Connect(ctx interface{}) *Client_Connect_Call {
	return &Client_Connect_Call{Call: _e.mock.On("Connect", ctx)}
}

func (_c *Client_Connect_Call) Run(run func(ctx context.Context)) *Client_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Client_Connect_Call) Return(_a0 error) *Client_Connect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Connect_Call) RunAndReturn(run func(context.Context) error) *Client_Connect_Call {
	_c.Call.Return(run)
	return _c
}

// Database provides a mock function with given fields: name, opts
func (_m *Client) Database(name string, opts ...*options.DatabaseOptions) mongoifc.Database {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 mongoifc.Database
	if rf, ok := ret.Get(0).(func(string, ...*options.DatabaseOptions) mongoifc.Database); ok {
		r0 = rf(name, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongoifc.Database)
		}
	}

	return r0
}

// Client_Database_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Database'
type Client_Database_Call struct {
	*mock.Call
}

// Database is a helper method to define mock.On call
//   - name string
//   - opts ...*options.DatabaseOptions
func (_e *Client_Expecter) Database(name interface{}, opts ...interface{}) *Client_Database_Call {
	return &Client_Database_Call{Call: _e.mock.On("Database",
		append([]interface{}{name}, opts...)...)}
}

func (_c *Client_Database_Call) Run(run func(name string, opts ...*options.DatabaseOptions)) *Client_Database_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*options.DatabaseOptions, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(*options.DatabaseOptions)
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Client_Database_Call) Return(_a0 mongoifc.Database) *Client_Database_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Database_Call) RunAndReturn(run func(string, ...*options.DatabaseOptions) mongoifc.Database) *Client_Database_Call {
	_c.Call.Return(run)
	return _c
}

// Disconnect provides a mock function with given fields: ctx
func (_m *Client) Disconnect(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Disconnect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Disconnect'
type Client_Disconnect_Call struct {
	*mock.Call
}

// Disconnect is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Client_Expecter) Disconnect(ctx interface{}) *Client_Disconnect_Call {
	return &Client_Disconnect_Call{Call: _e.mock.On("Disconnect", ctx)}
}

func (_c *Client_Disconnect_Call) Run(run func(ctx context.Context)) *Client_Disconnect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Client_Disconnect_Call) Return(_a0 error) *Client_Disconnect_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Disconnect_Call) RunAndReturn(run func(context.Context) error) *Client_Disconnect_Call {
	_c.Call.Return(run)
	return _c
}

// ListDatabaseNames provides a mock function with given fields: ctx, filter, opts
func (_m *Client) ListDatabaseNames(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) ([]string, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.ListDatabasesOptions) ([]string, error)); ok {
		return rf(ctx, filter, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.ListDatabasesOptions) []string); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.ListDatabasesOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_ListDatabaseNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDatabaseNames'
type Client_ListDatabaseNames_Call struct {
	*mock.Call
}

// ListDatabaseNames is a helper method to define mock.On call
//   - ctx context.Context
//   - filter interface{}
//   - opts ...*options.ListDatabasesOptions
func (_e *Client_Expecter) ListDatabaseNames(ctx interface{}, filter interface{}, opts ...interface{}) *Client_ListDatabaseNames_Call {
	return &Client_ListDatabaseNames_Call{Call: _e.mock.On("ListDatabaseNames",
		append([]interface{}{ctx, filter}, opts...)...)}
}

func (_c *Client_ListDatabaseNames_Call) Run(run func(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions)) *Client_ListDatabaseNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*options.ListDatabasesOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(*options.ListDatabasesOptions)
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *Client_ListDatabaseNames_Call) Return(_a0 []string, _a1 error) *Client_ListDatabaseNames_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_ListDatabaseNames_Call) RunAndReturn(run func(context.Context, interface{}, ...*options.ListDatabasesOptions) ([]string, error)) *Client_ListDatabaseNames_Call {
	_c.Call.Return(run)
	return _c
}

// ListDatabases provides a mock function with given fields: ctx, filter, opts
func (_m *Client) ListDatabases(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 mongo.ListDatabasesResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error)); ok {
		return rf(ctx, filter, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.ListDatabasesOptions) mongo.ListDatabasesResult); ok {
		r0 = rf(ctx, filter, opts...)
	} else {
		r0 = ret.Get(0).(mongo.ListDatabasesResult)
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.ListDatabasesOptions) error); ok {
		r1 = rf(ctx, filter, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_ListDatabases_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListDatabases'
type Client_ListDatabases_Call struct {
	*mock.Call
}

// ListDatabases is a helper method to define mock.On call
//   - ctx context.Context
//   - filter interface{}
//   - opts ...*options.ListDatabasesOptions
func (_e *Client_Expecter) ListDatabases(ctx interface{}, filter interface{}, opts ...interface{}) *Client_ListDatabases_Call {
	return &Client_ListDatabases_Call{Call: _e.mock.On("ListDatabases",
		append([]interface{}{ctx, filter}, opts...)...)}
}

func (_c *Client_ListDatabases_Call) Run(run func(ctx context.Context, filter interface{}, opts ...*options.ListDatabasesOptions)) *Client_ListDatabases_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*options.ListDatabasesOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(*options.ListDatabasesOptions)
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *Client_ListDatabases_Call) Return(_a0 mongo.ListDatabasesResult, _a1 error) *Client_ListDatabases_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_ListDatabases_Call) RunAndReturn(run func(context.Context, interface{}, ...*options.ListDatabasesOptions) (mongo.ListDatabasesResult, error)) *Client_ListDatabases_Call {
	_c.Call.Return(run)
	return _c
}

// NumberSessionsInProgress provides a mock function with given fields:
func (_m *Client) NumberSessionsInProgress() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Client_NumberSessionsInProgress_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NumberSessionsInProgress'
type Client_NumberSessionsInProgress_Call struct {
	*mock.Call
}

// NumberSessionsInProgress is a helper method to define mock.On call
func (_e *Client_Expecter) NumberSessionsInProgress() *Client_NumberSessionsInProgress_Call {
	return &Client_NumberSessionsInProgress_Call{Call: _e.mock.On("NumberSessionsInProgress")}
}

func (_c *Client_NumberSessionsInProgress_Call) Run(run func()) *Client_NumberSessionsInProgress_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_NumberSessionsInProgress_Call) Return(_a0 int) *Client_NumberSessionsInProgress_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_NumberSessionsInProgress_Call) RunAndReturn(run func() int) *Client_NumberSessionsInProgress_Call {
	_c.Call.Return(run)
	return _c
}

// Ping provides a mock function with given fields: ctx, rp
func (_m *Client) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	ret := _m.Called(ctx, rp)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *readpref.ReadPref) error); ok {
		r0 = rf(ctx, rp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_Ping_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ping'
type Client_Ping_Call struct {
	*mock.Call
}

// Ping is a helper method to define mock.On call
//   - ctx context.Context
//   - rp *readpref.ReadPref
func (_e *Client_Expecter) Ping(ctx interface{}, rp interface{}) *Client_Ping_Call {
	return &Client_Ping_Call{Call: _e.mock.On("Ping", ctx, rp)}
}

func (_c *Client_Ping_Call) Run(run func(ctx context.Context, rp *readpref.ReadPref)) *Client_Ping_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*readpref.ReadPref))
	})
	return _c
}

func (_c *Client_Ping_Call) Return(_a0 error) *Client_Ping_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Ping_Call) RunAndReturn(run func(context.Context, *readpref.ReadPref) error) *Client_Ping_Call {
	_c.Call.Return(run)
	return _c
}

// StartSession provides a mock function with given fields: opts
func (_m *Client) StartSession(opts ...*options.SessionOptions) (mongoifc.Session, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 mongoifc.Session
	var r1 error
	if rf, ok := ret.Get(0).(func(...*options.SessionOptions) (mongoifc.Session, error)); ok {
		return rf(opts...)
	}
	if rf, ok := ret.Get(0).(func(...*options.SessionOptions) mongoifc.Session); ok {
		r0 = rf(opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongoifc.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(...*options.SessionOptions) error); ok {
		r1 = rf(opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_StartSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartSession'
type Client_StartSession_Call struct {
	*mock.Call
}

// StartSession is a helper method to define mock.On call
//   - opts ...*options.SessionOptions
func (_e *Client_Expecter) StartSession(opts ...interface{}) *Client_StartSession_Call {
	return &Client_StartSession_Call{Call: _e.mock.On("StartSession",
		append([]interface{}{}, opts...)...)}
}

func (_c *Client_StartSession_Call) Run(run func(opts ...*options.SessionOptions)) *Client_StartSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*options.SessionOptions, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(*options.SessionOptions)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Client_StartSession_Call) Return(_a0 mongoifc.Session, _a1 error) *Client_StartSession_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_StartSession_Call) RunAndReturn(run func(...*options.SessionOptions) (mongoifc.Session, error)) *Client_StartSession_Call {
	_c.Call.Return(run)
	return _c
}

// Timeout provides a mock function with given fields:
func (_m *Client) Timeout() *time.Duration {
	ret := _m.Called()

	var r0 *time.Duration
	if rf, ok := ret.Get(0).(func() *time.Duration); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*time.Duration)
		}
	}

	return r0
}

// Client_Timeout_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Timeout'
type Client_Timeout_Call struct {
	*mock.Call
}

// Timeout is a helper method to define mock.On call
func (_e *Client_Expecter) Timeout() *Client_Timeout_Call {
	return &Client_Timeout_Call{Call: _e.mock.On("Timeout")}
}

func (_c *Client_Timeout_Call) Run(run func()) *Client_Timeout_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Client_Timeout_Call) Return(_a0 *time.Duration) *Client_Timeout_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_Timeout_Call) RunAndReturn(run func() *time.Duration) *Client_Timeout_Call {
	_c.Call.Return(run)
	return _c
}

// UseSession provides a mock function with given fields: ctx, fn
func (_m *Client) UseSession(ctx context.Context, fn func(mongoifc.SessionContext) error) error {
	ret := _m.Called(ctx, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(mongoifc.SessionContext) error) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_UseSession_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UseSession'
type Client_UseSession_Call struct {
	*mock.Call
}

// UseSession is a helper method to define mock.On call
//   - ctx context.Context
//   - fn func(mongoifc.SessionContext) error
func (_e *Client_Expecter) UseSession(ctx interface{}, fn interface{}) *Client_UseSession_Call {
	return &Client_UseSession_Call{Call: _e.mock.On("UseSession", ctx, fn)}
}

func (_c *Client_UseSession_Call) Run(run func(ctx context.Context, fn func(mongoifc.SessionContext) error)) *Client_UseSession_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(func(mongoifc.SessionContext) error))
	})
	return _c
}

func (_c *Client_UseSession_Call) Return(_a0 error) *Client_UseSession_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_UseSession_Call) RunAndReturn(run func(context.Context, func(mongoifc.SessionContext) error) error) *Client_UseSession_Call {
	_c.Call.Return(run)
	return _c
}

// UseSessionWithOptions provides a mock function with given fields: ctx, opts, fn
func (_m *Client) UseSessionWithOptions(ctx context.Context, opts *options.SessionOptions, fn func(mongoifc.SessionContext) error) error {
	ret := _m.Called(ctx, opts, fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *options.SessionOptions, func(mongoifc.SessionContext) error) error); ok {
		r0 = rf(ctx, opts, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Client_UseSessionWithOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UseSessionWithOptions'
type Client_UseSessionWithOptions_Call struct {
	*mock.Call
}

// UseSessionWithOptions is a helper method to define mock.On call
//   - ctx context.Context
//   - opts *options.SessionOptions
//   - fn func(mongoifc.SessionContext) error
func (_e *Client_Expecter) UseSessionWithOptions(ctx interface{}, opts interface{}, fn interface{}) *Client_UseSessionWithOptions_Call {
	return &Client_UseSessionWithOptions_Call{Call: _e.mock.On("UseSessionWithOptions", ctx, opts, fn)}
}

func (_c *Client_UseSessionWithOptions_Call) Run(run func(ctx context.Context, opts *options.SessionOptions, fn func(mongoifc.SessionContext) error)) *Client_UseSessionWithOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*options.SessionOptions), args[2].(func(mongoifc.SessionContext) error))
	})
	return _c
}

func (_c *Client_UseSessionWithOptions_Call) Return(_a0 error) *Client_UseSessionWithOptions_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Client_UseSessionWithOptions_Call) RunAndReturn(run func(context.Context, *options.SessionOptions, func(mongoifc.SessionContext) error) error) *Client_UseSessionWithOptions_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, pipeline, opts
func (_m *Client) Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (mongoifc.ChangeStream, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, pipeline)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 mongoifc.ChangeStream
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.ChangeStreamOptions) (mongoifc.ChangeStream, error)); ok {
		return rf(ctx, pipeline, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.ChangeStreamOptions) mongoifc.ChangeStream); ok {
		r0 = rf(ctx, pipeline, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongoifc.ChangeStream)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.ChangeStreamOptions) error); ok {
		r1 = rf(ctx, pipeline, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Client_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type Client_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - pipeline interface{}
//   - opts ...*options.ChangeStreamOptions
func (_e *Client_Expecter) Watch(ctx interface{}, pipeline interface{}, opts ...interface{}) *Client_Watch_Call {
	return &Client_Watch_Call{Call: _e.mock.On("Watch",
		append([]interface{}{ctx, pipeline}, opts...)...)}
}

func (_c *Client_Watch_Call) Run(run func(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions)) *Client_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*options.ChangeStreamOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(*options.ChangeStreamOptions)
			}
		}
		run(args[0].(context.Context), args[1].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *Client_Watch_Call) Return(_a0 mongoifc.ChangeStream, _a1 error) *Client_Watch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Client_Watch_Call) RunAndReturn(run func(context.Context, interface{}, ...*options.ChangeStreamOptions) (mongoifc.ChangeStream, error)) *Client_Watch_Call {
	_c.Call.Return(run)
	return _c
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
