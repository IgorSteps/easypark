// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// HandlerFactory is an autogenerated mock type for the HandlerFactory type
type HandlerFactory struct {
	mock.Mock
}

type HandlerFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *HandlerFactory) EXPECT() *HandlerFactory_Expecter {
	return &HandlerFactory_Expecter{mock: &_m.Mock}
}

// DriverBan provides a mock function with given fields:
func (_m *HandlerFactory) DriverBan() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DriverBan")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_DriverBan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DriverBan'
type HandlerFactory_DriverBan_Call struct {
	*mock.Call
}

// DriverBan is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) DriverBan() *HandlerFactory_DriverBan_Call {
	return &HandlerFactory_DriverBan_Call{Call: _e.mock.On("DriverBan")}
}

func (_c *HandlerFactory_DriverBan_Call) Run(run func()) *HandlerFactory_DriverBan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_DriverBan_Call) Return(_a0 http.Handler) *HandlerFactory_DriverBan_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_DriverBan_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_DriverBan_Call {
	_c.Call.Return(run)
	return _c
}

// DriverCreate provides a mock function with given fields:
func (_m *HandlerFactory) DriverCreate() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DriverCreate")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_DriverCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DriverCreate'
type HandlerFactory_DriverCreate_Call struct {
	*mock.Call
}

// DriverCreate is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) DriverCreate() *HandlerFactory_DriverCreate_Call {
	return &HandlerFactory_DriverCreate_Call{Call: _e.mock.On("DriverCreate")}
}

func (_c *HandlerFactory_DriverCreate_Call) Run(run func()) *HandlerFactory_DriverCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_DriverCreate_Call) Return(_a0 http.Handler) *HandlerFactory_DriverCreate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_DriverCreate_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_DriverCreate_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllDrivers provides a mock function with given fields:
func (_m *HandlerFactory) GetAllDrivers() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllDrivers")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_GetAllDrivers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllDrivers'
type HandlerFactory_GetAllDrivers_Call struct {
	*mock.Call
}

// GetAllDrivers is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) GetAllDrivers() *HandlerFactory_GetAllDrivers_Call {
	return &HandlerFactory_GetAllDrivers_Call{Call: _e.mock.On("GetAllDrivers")}
}

func (_c *HandlerFactory_GetAllDrivers_Call) Run(run func()) *HandlerFactory_GetAllDrivers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_GetAllDrivers_Call) Return(_a0 http.Handler) *HandlerFactory_GetAllDrivers_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_GetAllDrivers_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_GetAllDrivers_Call {
	_c.Call.Return(run)
	return _c
}

// UserAuthorise provides a mock function with given fields:
func (_m *HandlerFactory) UserAuthorise() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UserAuthorise")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_UserAuthorise_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserAuthorise'
type HandlerFactory_UserAuthorise_Call struct {
	*mock.Call
}

// UserAuthorise is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) UserAuthorise() *HandlerFactory_UserAuthorise_Call {
	return &HandlerFactory_UserAuthorise_Call{Call: _e.mock.On("UserAuthorise")}
}

func (_c *HandlerFactory_UserAuthorise_Call) Run(run func()) *HandlerFactory_UserAuthorise_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_UserAuthorise_Call) Return(_a0 http.Handler) *HandlerFactory_UserAuthorise_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_UserAuthorise_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_UserAuthorise_Call {
	_c.Call.Return(run)
	return _c
}

// NewHandlerFactory creates a new instance of HandlerFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandlerFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *HandlerFactory {
	mock := &HandlerFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
