// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	mongo "go.mongodb.org/mongo-driver/mongo"

	mongoifc "github.com/sv-tools/mongoifc"

	options "go.mongodb.org/mongo-driver/mongo/options"
)

// SearchIndexView is an autogenerated mock type for the SearchIndexView type
type SearchIndexView struct {
	mock.Mock
}

type SearchIndexView_Expecter struct {
	mock *mock.Mock
}

func (_m *SearchIndexView) EXPECT() *SearchIndexView_Expecter {
	return &SearchIndexView_Expecter{mock: &_m.Mock}
}

// CreateMany provides a mock function with given fields: ctx, models, opts
func (_m *SearchIndexView) CreateMany(ctx context.Context, models []mongo.SearchIndexModel, opts ...*options.CreateSearchIndexesOptions) ([]string, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, models)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateMany")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []mongo.SearchIndexModel, ...*options.CreateSearchIndexesOptions) ([]string, error)); ok {
		return rf(ctx, models, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []mongo.SearchIndexModel, ...*options.CreateSearchIndexesOptions) []string); ok {
		r0 = rf(ctx, models, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []mongo.SearchIndexModel, ...*options.CreateSearchIndexesOptions) error); ok {
		r1 = rf(ctx, models, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchIndexView_CreateMany_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMany'
type SearchIndexView_CreateMany_Call struct {
	*mock.Call
}

// CreateMany is a helper method to define mock.On call
//   - ctx context.Context
//   - models []mongo.SearchIndexModel
//   - opts ...*options.CreateSearchIndexesOptions
func (_e *SearchIndexView_Expecter) CreateMany(ctx interface{}, models interface{}, opts ...interface{}) *SearchIndexView_CreateMany_Call {
	return &SearchIndexView_CreateMany_Call{Call: _e.mock.On("CreateMany",
		append([]interface{}{ctx, models}, opts...)...)}
}

func (_c *SearchIndexView_CreateMany_Call) Run(run func(ctx context.Context, models []mongo.SearchIndexModel, opts ...*options.CreateSearchIndexesOptions)) *SearchIndexView_CreateMany_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*options.CreateSearchIndexesOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(*options.CreateSearchIndexesOptions)
			}
		}
		run(args[0].(context.Context), args[1].([]mongo.SearchIndexModel), variadicArgs...)
	})
	return _c
}

func (_c *SearchIndexView_CreateMany_Call) Return(_a0 []string, _a1 error) *SearchIndexView_CreateMany_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SearchIndexView_CreateMany_Call) RunAndReturn(run func(context.Context, []mongo.SearchIndexModel, ...*options.CreateSearchIndexesOptions) ([]string, error)) *SearchIndexView_CreateMany_Call {
	_c.Call.Return(run)
	return _c
}

// CreateOne provides a mock function with given fields: ctx, model, opts
func (_m *SearchIndexView) CreateOne(ctx context.Context, model mongo.SearchIndexModel, opts ...*options.CreateSearchIndexesOptions) (string, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, model)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateOne")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, mongo.SearchIndexModel, ...*options.CreateSearchIndexesOptions) (string, error)); ok {
		return rf(ctx, model, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, mongo.SearchIndexModel, ...*options.CreateSearchIndexesOptions) string); ok {
		r0 = rf(ctx, model, opts...)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, mongo.SearchIndexModel, ...*options.CreateSearchIndexesOptions) error); ok {
		r1 = rf(ctx, model, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchIndexView_CreateOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOne'
type SearchIndexView_CreateOne_Call struct {
	*mock.Call
}

// CreateOne is a helper method to define mock.On call
//   - ctx context.Context
//   - model mongo.SearchIndexModel
//   - opts ...*options.CreateSearchIndexesOptions
func (_e *SearchIndexView_Expecter) CreateOne(ctx interface{}, model interface{}, opts ...interface{}) *SearchIndexView_CreateOne_Call {
	return &SearchIndexView_CreateOne_Call{Call: _e.mock.On("CreateOne",
		append([]interface{}{ctx, model}, opts...)...)}
}

func (_c *SearchIndexView_CreateOne_Call) Run(run func(ctx context.Context, model mongo.SearchIndexModel, opts ...*options.CreateSearchIndexesOptions)) *SearchIndexView_CreateOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*options.CreateSearchIndexesOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(*options.CreateSearchIndexesOptions)
			}
		}
		run(args[0].(context.Context), args[1].(mongo.SearchIndexModel), variadicArgs...)
	})
	return _c
}

func (_c *SearchIndexView_CreateOne_Call) Return(_a0 string, _a1 error) *SearchIndexView_CreateOne_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SearchIndexView_CreateOne_Call) RunAndReturn(run func(context.Context, mongo.SearchIndexModel, ...*options.CreateSearchIndexesOptions) (string, error)) *SearchIndexView_CreateOne_Call {
	_c.Call.Return(run)
	return _c
}

// DropOne provides a mock function with given fields: ctx, name, opts
func (_m *SearchIndexView) DropOne(ctx context.Context, name string, opts ...*options.DropSearchIndexOptions) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DropOne")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, ...*options.DropSearchIndexOptions) error); ok {
		r0 = rf(ctx, name, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchIndexView_DropOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DropOne'
type SearchIndexView_DropOne_Call struct {
	*mock.Call
}

// DropOne is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - opts ...*options.DropSearchIndexOptions
func (_e *SearchIndexView_Expecter) DropOne(ctx interface{}, name interface{}, opts ...interface{}) *SearchIndexView_DropOne_Call {
	return &SearchIndexView_DropOne_Call{Call: _e.mock.On("DropOne",
		append([]interface{}{ctx, name}, opts...)...)}
}

func (_c *SearchIndexView_DropOne_Call) Run(run func(ctx context.Context, name string, opts ...*options.DropSearchIndexOptions)) *SearchIndexView_DropOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*options.DropSearchIndexOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(*options.DropSearchIndexOptions)
			}
		}
		run(args[0].(context.Context), args[1].(string), variadicArgs...)
	})
	return _c
}

func (_c *SearchIndexView_DropOne_Call) Return(_a0 error) *SearchIndexView_DropOne_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SearchIndexView_DropOne_Call) RunAndReturn(run func(context.Context, string, ...*options.DropSearchIndexOptions) error) *SearchIndexView_DropOne_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, searchIdxOpts, opts
func (_m *SearchIndexView) List(ctx context.Context, searchIdxOpts *options.SearchIndexesOptions, opts ...*options.ListSearchIndexesOptions) (mongoifc.Cursor, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, searchIdxOpts)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 mongoifc.Cursor
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *options.SearchIndexesOptions, ...*options.ListSearchIndexesOptions) (mongoifc.Cursor, error)); ok {
		return rf(ctx, searchIdxOpts, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *options.SearchIndexesOptions, ...*options.ListSearchIndexesOptions) mongoifc.Cursor); ok {
		r0 = rf(ctx, searchIdxOpts, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(mongoifc.Cursor)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *options.SearchIndexesOptions, ...*options.ListSearchIndexesOptions) error); ok {
		r1 = rf(ctx, searchIdxOpts, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchIndexView_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type SearchIndexView_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - searchIdxOpts *options.SearchIndexesOptions
//   - opts ...*options.ListSearchIndexesOptions
func (_e *SearchIndexView_Expecter) List(ctx interface{}, searchIdxOpts interface{}, opts ...interface{}) *SearchIndexView_List_Call {
	return &SearchIndexView_List_Call{Call: _e.mock.On("List",
		append([]interface{}{ctx, searchIdxOpts}, opts...)...)}
}

func (_c *SearchIndexView_List_Call) Run(run func(ctx context.Context, searchIdxOpts *options.SearchIndexesOptions, opts ...*options.ListSearchIndexesOptions)) *SearchIndexView_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*options.ListSearchIndexesOptions, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(*options.ListSearchIndexesOptions)
			}
		}
		run(args[0].(context.Context), args[1].(*options.SearchIndexesOptions), variadicArgs...)
	})
	return _c
}

func (_c *SearchIndexView_List_Call) Return(_a0 mongoifc.Cursor, _a1 error) *SearchIndexView_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *SearchIndexView_List_Call) RunAndReturn(run func(context.Context, *options.SearchIndexesOptions, ...*options.ListSearchIndexesOptions) (mongoifc.Cursor, error)) *SearchIndexView_List_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateOne provides a mock function with given fields: ctx, name, definition, opts
func (_m *SearchIndexView) UpdateOne(ctx context.Context, name string, definition interface{}, opts ...*options.UpdateSearchIndexOptions) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, name, definition)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOne")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, ...*options.UpdateSearchIndexOptions) error); ok {
		r0 = rf(ctx, name, definition, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchIndexView_UpdateOne_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateOne'
type SearchIndexView_UpdateOne_Call struct {
	*mock.Call
}

// UpdateOne is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - definition interface{}
//   - opts ...*options.UpdateSearchIndexOptions
func (_e *SearchIndexView_Expecter) UpdateOne(ctx interface{}, name interface{}, definition interface{}, opts ...interface{}) *SearchIndexView_UpdateOne_Call {
	return &SearchIndexView_UpdateOne_Call{Call: _e.mock.On("UpdateOne",
		append([]interface{}{ctx, name, definition}, opts...)...)}
}

func (_c *SearchIndexView_UpdateOne_Call) Run(run func(ctx context.Context, name string, definition interface{}, opts ...*options.UpdateSearchIndexOptions)) *SearchIndexView_UpdateOne_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*options.UpdateSearchIndexOptions, len(args)-3)
		for i, a := range args[3:] {
			if a != nil {
				variadicArgs[i] = a.(*options.UpdateSearchIndexOptions)
			}
		}
		run(args[0].(context.Context), args[1].(string), args[2].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *SearchIndexView_UpdateOne_Call) Return(_a0 error) *SearchIndexView_UpdateOne_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *SearchIndexView_UpdateOne_Call) RunAndReturn(run func(context.Context, string, interface{}, ...*options.UpdateSearchIndexOptions) error) *SearchIndexView_UpdateOne_Call {
	_c.Call.Return(run)
	return _c
}

// NewSearchIndexView creates a new instance of SearchIndexView. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSearchIndexView(t interface {
	mock.TestingT
	Cleanup(func())
}) *SearchIndexView {
	mock := &SearchIndexView{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
