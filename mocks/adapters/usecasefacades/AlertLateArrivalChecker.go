// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// AlertLateArrivalChecker is an autogenerated mock type for the AlertLateArrivalChecker type
type AlertLateArrivalChecker struct {
	mock.Mock
}

type AlertLateArrivalChecker_Expecter struct {
	mock *mock.Mock
}

func (_m *AlertLateArrivalChecker) EXPECT() *AlertLateArrivalChecker_Expecter {
	return &AlertLateArrivalChecker_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, threshold
func (_m *AlertLateArrivalChecker) Execute(ctx context.Context, threshold time.Duration) ([]entities.Alert, error) {
	ret := _m.Called(ctx, threshold)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 []entities.Alert
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration) ([]entities.Alert, error)); ok {
		return rf(ctx, threshold)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Duration) []entities.Alert); ok {
		r0 = rf(ctx, threshold)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Alert)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Duration) error); ok {
		r1 = rf(ctx, threshold)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AlertLateArrivalChecker_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type AlertLateArrivalChecker_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - threshold time.Duration
func (_e *AlertLateArrivalChecker_Expecter) Execute(ctx interface{}, threshold interface{}) *AlertLateArrivalChecker_Execute_Call {
	return &AlertLateArrivalChecker_Execute_Call{Call: _e.mock.On("Execute", ctx, threshold)}
}

func (_c *AlertLateArrivalChecker_Execute_Call) Run(run func(ctx context.Context, threshold time.Duration)) *AlertLateArrivalChecker_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(time.Duration))
	})
	return _c
}

func (_c *AlertLateArrivalChecker_Execute_Call) Return(_a0 []entities.Alert, _a1 error) *AlertLateArrivalChecker_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AlertLateArrivalChecker_Execute_Call) RunAndReturn(run func(context.Context, time.Duration) ([]entities.Alert, error)) *AlertLateArrivalChecker_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewAlertLateArrivalChecker creates a new instance of AlertLateArrivalChecker. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAlertLateArrivalChecker(t interface {
	mock.TestingT
	Cleanup(func())
}) *AlertLateArrivalChecker {
	mock := &AlertLateArrivalChecker{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
