// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	datastore "github.com/IgorSteps/easypark/internal/adapters/datastore"
	mock "github.com/stretchr/testify/mock"
)

// Datastore is an autogenerated mock type for the Datastore type
type Datastore struct {
	mock.Mock
}

type Datastore_Expecter struct {
	mock *mock.Mock
}

func (_m *Datastore) EXPECT() *Datastore_Expecter {
	return &Datastore_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: value
func (_m *Datastore) Create(value interface{}) datastore.Datastore {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 datastore.Datastore
	if rf, ok := ret.Get(0).(func(interface{}) datastore.Datastore); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.Datastore)
		}
	}

	return r0
}

// Datastore_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Datastore_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - value interface{}
func (_e *Datastore_Expecter) Create(value interface{}) *Datastore_Create_Call {
	return &Datastore_Create_Call{Call: _e.mock.On("Create", value)}
}

func (_c *Datastore_Create_Call) Run(run func(value interface{})) *Datastore_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *Datastore_Create_Call) Return(_a0 datastore.Datastore) *Datastore_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Datastore_Create_Call) RunAndReturn(run func(interface{}) datastore.Datastore) *Datastore_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: value, conds
func (_m *Datastore) Delete(value interface{}, conds ...interface{}) datastore.Datastore {
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, conds...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 datastore.Datastore
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) datastore.Datastore); ok {
		r0 = rf(value, conds...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.Datastore)
		}
	}

	return r0
}

// Datastore_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type Datastore_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - value interface{}
//   - conds ...interface{}
func (_e *Datastore_Expecter) Delete(value interface{}, conds ...interface{}) *Datastore_Delete_Call {
	return &Datastore_Delete_Call{Call: _e.mock.On("Delete",
		append([]interface{}{value}, conds...)...)}
}

func (_c *Datastore_Delete_Call) Run(run func(value interface{}, conds ...interface{})) *Datastore_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *Datastore_Delete_Call) Return(_a0 datastore.Datastore) *Datastore_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Datastore_Delete_Call) RunAndReturn(run func(interface{}, ...interface{}) datastore.Datastore) *Datastore_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Error provides a mock function with given fields:
func (_m *Datastore) Error() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Error")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Datastore_Error_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Error'
type Datastore_Error_Call struct {
	*mock.Call
}

// Error is a helper method to define mock.On call
func (_e *Datastore_Expecter) Error() *Datastore_Error_Call {
	return &Datastore_Error_Call{Call: _e.mock.On("Error")}
}

func (_c *Datastore_Error_Call) Run(run func()) *Datastore_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Datastore_Error_Call) Return(_a0 error) *Datastore_Error_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Datastore_Error_Call) RunAndReturn(run func() error) *Datastore_Error_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields: value
func (_m *Datastore) FindAll(value interface{}) datastore.Datastore {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 datastore.Datastore
	if rf, ok := ret.Get(0).(func(interface{}) datastore.Datastore); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.Datastore)
		}
	}

	return r0
}

// Datastore_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type Datastore_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
//   - value interface{}
func (_e *Datastore_Expecter) FindAll(value interface{}) *Datastore_FindAll_Call {
	return &Datastore_FindAll_Call{Call: _e.mock.On("FindAll", value)}
}

func (_c *Datastore_FindAll_Call) Run(run func(value interface{})) *Datastore_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *Datastore_FindAll_Call) Return(_a0 datastore.Datastore) *Datastore_FindAll_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Datastore_FindAll_Call) RunAndReturn(run func(interface{}) datastore.Datastore) *Datastore_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// First provides a mock function with given fields: value, args
func (_m *Datastore) First(value interface{}, args ...interface{}) datastore.Datastore {
	var _ca []interface{}
	_ca = append(_ca, value)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for First")
	}

	var r0 datastore.Datastore
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) datastore.Datastore); ok {
		r0 = rf(value, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.Datastore)
		}
	}

	return r0
}

// Datastore_First_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'First'
type Datastore_First_Call struct {
	*mock.Call
}

// First is a helper method to define mock.On call
//   - value interface{}
//   - args ...interface{}
func (_e *Datastore_Expecter) First(value interface{}, args ...interface{}) *Datastore_First_Call {
	return &Datastore_First_Call{Call: _e.mock.On("First",
		append([]interface{}{value}, args...)...)}
}

func (_c *Datastore_First_Call) Run(run func(value interface{}, args ...interface{})) *Datastore_First_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *Datastore_First_Call) Return(_a0 datastore.Datastore) *Datastore_First_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Datastore_First_Call) RunAndReturn(run func(interface{}, ...interface{}) datastore.Datastore) *Datastore_First_Call {
	_c.Call.Return(run)
	return _c
}

// Preload provides a mock function with given fields: column, conditions
func (_m *Datastore) Preload(column string, conditions ...interface{}) datastore.Datastore {
	var _ca []interface{}
	_ca = append(_ca, column)
	_ca = append(_ca, conditions...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Preload")
	}

	var r0 datastore.Datastore
	if rf, ok := ret.Get(0).(func(string, ...interface{}) datastore.Datastore); ok {
		r0 = rf(column, conditions...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.Datastore)
		}
	}

	return r0
}

// Datastore_Preload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Preload'
type Datastore_Preload_Call struct {
	*mock.Call
}

// Preload is a helper method to define mock.On call
//   - column string
//   - conditions ...interface{}
func (_e *Datastore_Expecter) Preload(column interface{}, conditions ...interface{}) *Datastore_Preload_Call {
	return &Datastore_Preload_Call{Call: _e.mock.On("Preload",
		append([]interface{}{column}, conditions...)...)}
}

func (_c *Datastore_Preload_Call) Run(run func(column string, conditions ...interface{})) *Datastore_Preload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), variadicArgs...)
	})
	return _c
}

func (_c *Datastore_Preload_Call) Return(_a0 datastore.Datastore) *Datastore_Preload_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Datastore_Preload_Call) RunAndReturn(run func(string, ...interface{}) datastore.Datastore) *Datastore_Preload_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: value
func (_m *Datastore) Save(value interface{}) datastore.Datastore {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 datastore.Datastore
	if rf, ok := ret.Get(0).(func(interface{}) datastore.Datastore); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.Datastore)
		}
	}

	return r0
}

// Datastore_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type Datastore_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - value interface{}
func (_e *Datastore_Expecter) Save(value interface{}) *Datastore_Save_Call {
	return &Datastore_Save_Call{Call: _e.mock.On("Save", value)}
}

func (_c *Datastore_Save_Call) Run(run func(value interface{})) *Datastore_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *Datastore_Save_Call) Return(_a0 datastore.Datastore) *Datastore_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Datastore_Save_Call) RunAndReturn(run func(interface{}) datastore.Datastore) *Datastore_Save_Call {
	_c.Call.Return(run)
	return _c
}

// Where provides a mock function with given fields: query, args
func (_m *Datastore) Where(query interface{}, args ...interface{}) datastore.Datastore {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Where")
	}

	var r0 datastore.Datastore
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) datastore.Datastore); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.Datastore)
		}
	}

	return r0
}

// Datastore_Where_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Where'
type Datastore_Where_Call struct {
	*mock.Call
}

// Where is a helper method to define mock.On call
//   - query interface{}
//   - args ...interface{}
func (_e *Datastore_Expecter) Where(query interface{}, args ...interface{}) *Datastore_Where_Call {
	return &Datastore_Where_Call{Call: _e.mock.On("Where",
		append([]interface{}{query}, args...)...)}
}

func (_c *Datastore_Where_Call) Run(run func(query interface{}, args ...interface{})) *Datastore_Where_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(interface{}), variadicArgs...)
	})
	return _c
}

func (_c *Datastore_Where_Call) Return(_a0 datastore.Datastore) *Datastore_Where_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Datastore_Where_Call) RunAndReturn(run func(interface{}, ...interface{}) datastore.Datastore) *Datastore_Where_Call {
	_c.Call.Return(run)
	return _c
}

// WithContext provides a mock function with given fields: ctx
func (_m *Datastore) WithContext(ctx context.Context) datastore.Datastore {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for WithContext")
	}

	var r0 datastore.Datastore
	if rf, ok := ret.Get(0).(func(context.Context) datastore.Datastore); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(datastore.Datastore)
		}
	}

	return r0
}

// Datastore_WithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithContext'
type Datastore_WithContext_Call struct {
	*mock.Call
}

// WithContext is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Datastore_Expecter) WithContext(ctx interface{}) *Datastore_WithContext_Call {
	return &Datastore_WithContext_Call{Call: _e.mock.On("WithContext", ctx)}
}

func (_c *Datastore_WithContext_Call) Run(run func(ctx context.Context)) *Datastore_WithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Datastore_WithContext_Call) Return(_a0 datastore.Datastore) *Datastore_WithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Datastore_WithContext_Call) RunAndReturn(run func(context.Context) datastore.Datastore) *Datastore_WithContext_Call {
	_c.Call.Return(run)
	return _c
}

// NewDatastore creates a new instance of Datastore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDatastore(t interface {
	mock.TestingT
	Cleanup(func())
}) *Datastore {
	mock := &Datastore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
