// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// AlertSingleGetter is an autogenerated mock type for the AlertSingleGetter type
type AlertSingleGetter struct {
	mock.Mock
}

type AlertSingleGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *AlertSingleGetter) EXPECT() *AlertSingleGetter_Expecter {
	return &AlertSingleGetter_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, id
func (_m *AlertSingleGetter) Execute(ctx context.Context, id uuid.UUID) (entities.Alert, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 entities.Alert
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (entities.Alert, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) entities.Alert); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entities.Alert)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AlertSingleGetter_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type AlertSingleGetter_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *AlertSingleGetter_Expecter) Execute(ctx interface{}, id interface{}) *AlertSingleGetter_Execute_Call {
	return &AlertSingleGetter_Execute_Call{Call: _e.mock.On("Execute", ctx, id)}
}

func (_c *AlertSingleGetter_Execute_Call) Run(run func(ctx context.Context, id uuid.UUID)) *AlertSingleGetter_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *AlertSingleGetter_Execute_Call) Return(_a0 entities.Alert, _a1 error) *AlertSingleGetter_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AlertSingleGetter_Execute_Call) RunAndReturn(run func(context.Context, uuid.UUID) (entities.Alert, error)) *AlertSingleGetter_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewAlertSingleGetter creates a new instance of AlertSingleGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAlertSingleGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *AlertSingleGetter {
	mock := &AlertSingleGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
