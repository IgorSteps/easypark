// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"
	mock "github.com/stretchr/testify/mock"
)

// UserAuthenticator is an autogenerated mock type for the UserAuthenticator type
type UserAuthenticator struct {
	mock.Mock
}

type UserAuthenticator_Expecter struct {
	mock *mock.Mock
}

func (_m *UserAuthenticator) EXPECT() *UserAuthenticator_Expecter {
	return &UserAuthenticator_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, username, password
func (_m *UserAuthenticator) Execute(ctx context.Context, username string, password string) (*entities.User, string, error) {
	ret := _m.Called(ctx, username, password)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *entities.User
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*entities.User, string, error)); ok {
		return rf(ctx, username, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *entities.User); ok {
		r0 = rf(ctx, username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) string); ok {
		r1 = rf(ctx, username, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string, string) error); ok {
		r2 = rf(ctx, username, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UserAuthenticator_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type UserAuthenticator_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
//   - password string
func (_e *UserAuthenticator_Expecter) Execute(ctx interface{}, username interface{}, password interface{}) *UserAuthenticator_Execute_Call {
	return &UserAuthenticator_Execute_Call{Call: _e.mock.On("Execute", ctx, username, password)}
}

func (_c *UserAuthenticator_Execute_Call) Run(run func(ctx context.Context, username string, password string)) *UserAuthenticator_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *UserAuthenticator_Execute_Call) Return(_a0 *entities.User, _a1 string, _a2 error) *UserAuthenticator_Execute_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *UserAuthenticator_Execute_Call) RunAndReturn(run func(context.Context, string, string) (*entities.User, string, error)) *UserAuthenticator_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserAuthenticator creates a new instance of UserAuthenticator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserAuthenticator(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserAuthenticator {
	mock := &UserAuthenticator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
