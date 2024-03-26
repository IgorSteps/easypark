// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// Middleware is an autogenerated mock type for the Middleware type
type Middleware struct {
	mock.Mock
}

type Middleware_Expecter struct {
	mock *mock.Mock
}

func (_m *Middleware) EXPECT() *Middleware_Expecter {
	return &Middleware_Expecter{mock: &_m.Mock}
}

// Authorise provides a mock function with given fields: next
func (_m *Middleware) Authorise(next http.Handler) http.Handler {
	ret := _m.Called(next)

	if len(ret) == 0 {
		panic("no return value specified for Authorise")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(http.Handler) http.Handler); ok {
		r0 = rf(next)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// Middleware_Authorise_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Authorise'
type Middleware_Authorise_Call struct {
	*mock.Call
}

// Authorise is a helper method to define mock.On call
//   - next http.Handler
func (_e *Middleware_Expecter) Authorise(next interface{}) *Middleware_Authorise_Call {
	return &Middleware_Authorise_Call{Call: _e.mock.On("Authorise", next)}
}

func (_c *Middleware_Authorise_Call) Run(run func(next http.Handler)) *Middleware_Authorise_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.Handler))
	})
	return _c
}

func (_c *Middleware_Authorise_Call) Return(_a0 http.Handler) *Middleware_Authorise_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Middleware_Authorise_Call) RunAndReturn(run func(http.Handler) http.Handler) *Middleware_Authorise_Call {
	_c.Call.Return(run)
	return _c
}

// CheckStatus provides a mock function with given fields: next
func (_m *Middleware) CheckStatus(next http.Handler) http.Handler {
	ret := _m.Called(next)

	if len(ret) == 0 {
		panic("no return value specified for CheckStatus")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func(http.Handler) http.Handler); ok {
		r0 = rf(next)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// Middleware_CheckStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckStatus'
type Middleware_CheckStatus_Call struct {
	*mock.Call
}

// CheckStatus is a helper method to define mock.On call
//   - next http.Handler
func (_e *Middleware_Expecter) CheckStatus(next interface{}) *Middleware_CheckStatus_Call {
	return &Middleware_CheckStatus_Call{Call: _e.mock.On("CheckStatus", next)}
}

func (_c *Middleware_CheckStatus_Call) Run(run func(next http.Handler)) *Middleware_CheckStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.Handler))
	})
	return _c
}

func (_c *Middleware_CheckStatus_Call) Return(_a0 http.Handler) *Middleware_CheckStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Middleware_CheckStatus_Call) RunAndReturn(run func(http.Handler) http.Handler) *Middleware_CheckStatus_Call {
	_c.Call.Return(run)
	return _c
}

// RequireRole provides a mock function with given fields: requiredRole
func (_m *Middleware) RequireRole(requiredRole entities.UserRole) func(http.Handler) http.Handler {
	ret := _m.Called(requiredRole)

	if len(ret) == 0 {
		panic("no return value specified for RequireRole")
	}

	var r0 func(http.Handler) http.Handler
	if rf, ok := ret.Get(0).(func(entities.UserRole) func(http.Handler) http.Handler); ok {
		r0 = rf(requiredRole)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(http.Handler) http.Handler)
		}
	}

	return r0
}

// Middleware_RequireRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RequireRole'
type Middleware_RequireRole_Call struct {
	*mock.Call
}

// RequireRole is a helper method to define mock.On call
//   - requiredRole entities.UserRole
func (_e *Middleware_Expecter) RequireRole(requiredRole interface{}) *Middleware_RequireRole_Call {
	return &Middleware_RequireRole_Call{Call: _e.mock.On("RequireRole", requiredRole)}
}

func (_c *Middleware_RequireRole_Call) Run(run func(requiredRole entities.UserRole)) *Middleware_RequireRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entities.UserRole))
	})
	return _c
}

func (_c *Middleware_RequireRole_Call) Return(_a0 func(http.Handler) http.Handler) *Middleware_RequireRole_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Middleware_RequireRole_Call) RunAndReturn(run func(entities.UserRole) func(http.Handler) http.Handler) *Middleware_RequireRole_Call {
	_c.Call.Return(run)
	return _c
}

// NewMiddleware creates a new instance of Middleware. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMiddleware(t interface {
	mock.TestingT
	Cleanup(func())
}) *Middleware {
	mock := &Middleware{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
