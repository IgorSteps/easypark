// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"
	mock "github.com/stretchr/testify/mock"
)

// ParkingRequestsGetter is an autogenerated mock type for the ParkingRequestsGetter type
type ParkingRequestsGetter struct {
	mock.Mock
}

type ParkingRequestsGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *ParkingRequestsGetter) EXPECT() *ParkingRequestsGetter_Expecter {
	return &ParkingRequestsGetter_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx
func (_m *ParkingRequestsGetter) Execute(ctx context.Context) ([]entities.ParkingRequest, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 []entities.ParkingRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entities.ParkingRequest, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entities.ParkingRequest); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.ParkingRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParkingRequestsGetter_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type ParkingRequestsGetter_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ParkingRequestsGetter_Expecter) Execute(ctx interface{}) *ParkingRequestsGetter_Execute_Call {
	return &ParkingRequestsGetter_Execute_Call{Call: _e.mock.On("Execute", ctx)}
}

func (_c *ParkingRequestsGetter_Execute_Call) Run(run func(ctx context.Context)) *ParkingRequestsGetter_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ParkingRequestsGetter_Execute_Call) Return(_a0 []entities.ParkingRequest, _a1 error) *ParkingRequestsGetter_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ParkingRequestsGetter_Execute_Call) RunAndReturn(run func(context.Context) ([]entities.ParkingRequest, error)) *ParkingRequestsGetter_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewParkingRequestsGetter creates a new instance of ParkingRequestsGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewParkingRequestsGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *ParkingRequestsGetter {
	mock := &ParkingRequestsGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
