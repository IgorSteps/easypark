// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"
	mock "github.com/stretchr/testify/mock"
)

// DriverCreator is an autogenerated mock type for the DriverCreator type
type DriverCreator struct {
	mock.Mock
}

type DriverCreator_Expecter struct {
	mock *mock.Mock
}

func (_m *DriverCreator) EXPECT() *DriverCreator_Expecter {
	return &DriverCreator_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, user
func (_m *DriverCreator) Execute(ctx context.Context, user *entities.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DriverCreator_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type DriverCreator_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - user *entities.User
func (_e *DriverCreator_Expecter) Execute(ctx interface{}, user interface{}) *DriverCreator_Execute_Call {
	return &DriverCreator_Execute_Call{Call: _e.mock.On("Execute", ctx, user)}
}

func (_c *DriverCreator_Execute_Call) Run(run func(ctx context.Context, user *entities.User)) *DriverCreator_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.User))
	})
	return _c
}

func (_c *DriverCreator_Execute_Call) Return(_a0 error) *DriverCreator_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DriverCreator_Execute_Call) RunAndReturn(run func(context.Context, *entities.User) error) *DriverCreator_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewDriverCreator creates a new instance of DriverCreator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDriverCreator(t interface {
	mock.TestingT
	Cleanup(func())
}) *DriverCreator {
	mock := &DriverCreator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
