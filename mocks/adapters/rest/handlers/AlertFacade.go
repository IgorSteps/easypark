// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"

	mock "github.com/stretchr/testify/mock"

	time "time"

	uuid "github.com/google/uuid"
)

// AlertFacade is an autogenerated mock type for the AlertFacade type
type AlertFacade struct {
	mock.Mock
}

type AlertFacade_Expecter struct {
	mock *mock.Mock
}

func (_m *AlertFacade) EXPECT() *AlertFacade_Expecter {
	return &AlertFacade_Expecter{mock: &_m.Mock}
}

// CheckForLateArrivals provides a mock function with given fields: ctx, threshold
func (_m *AlertFacade) CheckForLateArrivals(ctx context.Context, threshold time.Duration) ([]entities.Alert, error) {
	ret := _m.Called(ctx, threshold)

	if len(ret) == 0 {
		panic("no return value specified for CheckForLateArrivals")
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

// AlertFacade_CheckForLateArrivals_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckForLateArrivals'
type AlertFacade_CheckForLateArrivals_Call struct {
	*mock.Call
}

// CheckForLateArrivals is a helper method to define mock.On call
//   - ctx context.Context
//   - threshold time.Duration
func (_e *AlertFacade_Expecter) CheckForLateArrivals(ctx interface{}, threshold interface{}) *AlertFacade_CheckForLateArrivals_Call {
	return &AlertFacade_CheckForLateArrivals_Call{Call: _e.mock.On("CheckForLateArrivals", ctx, threshold)}
}

func (_c *AlertFacade_CheckForLateArrivals_Call) Run(run func(ctx context.Context, threshold time.Duration)) *AlertFacade_CheckForLateArrivals_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(time.Duration))
	})
	return _c
}

func (_c *AlertFacade_CheckForLateArrivals_Call) Return(_a0 []entities.Alert, _a1 error) *AlertFacade_CheckForLateArrivals_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AlertFacade_CheckForLateArrivals_Call) RunAndReturn(run func(context.Context, time.Duration) ([]entities.Alert, error)) *AlertFacade_CheckForLateArrivals_Call {
	_c.Call.Return(run)
	return _c
}

// CheckForOverStays provides a mock function with given fields: ctx, threshold
func (_m *AlertFacade) CheckForOverStays(ctx context.Context, threshold time.Duration) ([]entities.Alert, error) {
	ret := _m.Called(ctx, threshold)

	if len(ret) == 0 {
		panic("no return value specified for CheckForOverStays")
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

// AlertFacade_CheckForOverStays_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckForOverStays'
type AlertFacade_CheckForOverStays_Call struct {
	*mock.Call
}

// CheckForOverStays is a helper method to define mock.On call
//   - ctx context.Context
//   - threshold time.Duration
func (_e *AlertFacade_Expecter) CheckForOverStays(ctx interface{}, threshold interface{}) *AlertFacade_CheckForOverStays_Call {
	return &AlertFacade_CheckForOverStays_Call{Call: _e.mock.On("CheckForOverStays", ctx, threshold)}
}

func (_c *AlertFacade_CheckForOverStays_Call) Run(run func(ctx context.Context, threshold time.Duration)) *AlertFacade_CheckForOverStays_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(time.Duration))
	})
	return _c
}

func (_c *AlertFacade_CheckForOverStays_Call) Return(_a0 []entities.Alert, _a1 error) *AlertFacade_CheckForOverStays_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AlertFacade_CheckForOverStays_Call) RunAndReturn(run func(context.Context, time.Duration) ([]entities.Alert, error)) *AlertFacade_CheckForOverStays_Call {
	_c.Call.Return(run)
	return _c
}

// GetAlert provides a mock function with given fields: ctx, id
func (_m *AlertFacade) GetAlert(ctx context.Context, id uuid.UUID) (entities.Alert, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetAlert")
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

// AlertFacade_GetAlert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAlert'
type AlertFacade_GetAlert_Call struct {
	*mock.Call
}

// GetAlert is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *AlertFacade_Expecter) GetAlert(ctx interface{}, id interface{}) *AlertFacade_GetAlert_Call {
	return &AlertFacade_GetAlert_Call{Call: _e.mock.On("GetAlert", ctx, id)}
}

func (_c *AlertFacade_GetAlert_Call) Run(run func(ctx context.Context, id uuid.UUID)) *AlertFacade_GetAlert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *AlertFacade_GetAlert_Call) Return(_a0 entities.Alert, _a1 error) *AlertFacade_GetAlert_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AlertFacade_GetAlert_Call) RunAndReturn(run func(context.Context, uuid.UUID) (entities.Alert, error)) *AlertFacade_GetAlert_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllAlerts provides a mock function with given fields: ctx
func (_m *AlertFacade) GetAllAlerts(ctx context.Context) ([]entities.Alert, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllAlerts")
	}

	var r0 []entities.Alert
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entities.Alert, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entities.Alert); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Alert)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AlertFacade_GetAllAlerts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllAlerts'
type AlertFacade_GetAllAlerts_Call struct {
	*mock.Call
}

// GetAllAlerts is a helper method to define mock.On call
//   - ctx context.Context
func (_e *AlertFacade_Expecter) GetAllAlerts(ctx interface{}) *AlertFacade_GetAllAlerts_Call {
	return &AlertFacade_GetAllAlerts_Call{Call: _e.mock.On("GetAllAlerts", ctx)}
}

func (_c *AlertFacade_GetAllAlerts_Call) Run(run func(ctx context.Context)) *AlertFacade_GetAllAlerts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *AlertFacade_GetAllAlerts_Call) Return(_a0 []entities.Alert, _a1 error) *AlertFacade_GetAllAlerts_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AlertFacade_GetAllAlerts_Call) RunAndReturn(run func(context.Context) ([]entities.Alert, error)) *AlertFacade_GetAllAlerts_Call {
	_c.Call.Return(run)
	return _c
}

// NewAlertFacade creates a new instance of AlertFacade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAlertFacade(t interface {
	mock.TestingT
	Cleanup(func())
}) *AlertFacade {
	mock := &AlertFacade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
