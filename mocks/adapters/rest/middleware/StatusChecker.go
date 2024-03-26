// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// StatusChecker is an autogenerated mock type for the StatusChecker type
type StatusChecker struct {
	mock.Mock
}

type StatusChecker_Expecter struct {
	mock *mock.Mock
}

func (_m *StatusChecker) EXPECT() *StatusChecker_Expecter {
	return &StatusChecker_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, id
func (_m *StatusChecker) Execute(ctx context.Context, id uuid.UUID) (bool, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (bool, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) bool); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatusChecker_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type StatusChecker_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *StatusChecker_Expecter) Execute(ctx interface{}, id interface{}) *StatusChecker_Execute_Call {
	return &StatusChecker_Execute_Call{Call: _e.mock.On("Execute", ctx, id)}
}

func (_c *StatusChecker_Execute_Call) Run(run func(ctx context.Context, id uuid.UUID)) *StatusChecker_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *StatusChecker_Execute_Call) Return(_a0 bool, _a1 error) *StatusChecker_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StatusChecker_Execute_Call) RunAndReturn(run func(context.Context, uuid.UUID) (bool, error)) *StatusChecker_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewStatusChecker creates a new instance of StatusChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStatusChecker(t interface {
	mock.TestingT
	Cleanup(func())
}) *StatusChecker {
	mock := &StatusChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}