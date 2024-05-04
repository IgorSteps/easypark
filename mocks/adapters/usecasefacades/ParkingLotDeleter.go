// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ParkingLotDeleter is an autogenerated mock type for the ParkingLotDeleter type
type ParkingLotDeleter struct {
	mock.Mock
}

type ParkingLotDeleter_Expecter struct {
	mock *mock.Mock
}

func (_m *ParkingLotDeleter) EXPECT() *ParkingLotDeleter_Expecter {
	return &ParkingLotDeleter_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, id
func (_m *ParkingLotDeleter) Execute(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ParkingLotDeleter_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type ParkingLotDeleter_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *ParkingLotDeleter_Expecter) Execute(ctx interface{}, id interface{}) *ParkingLotDeleter_Execute_Call {
	return &ParkingLotDeleter_Execute_Call{Call: _e.mock.On("Execute", ctx, id)}
}

func (_c *ParkingLotDeleter_Execute_Call) Run(run func(ctx context.Context, id uuid.UUID)) *ParkingLotDeleter_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *ParkingLotDeleter_Execute_Call) Return(_a0 error) *ParkingLotDeleter_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ParkingLotDeleter_Execute_Call) RunAndReturn(run func(context.Context, uuid.UUID) error) *ParkingLotDeleter_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewParkingLotDeleter creates a new instance of ParkingLotDeleter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewParkingLotDeleter(t interface {
	mock.TestingT
	Cleanup(func())
}) *ParkingLotDeleter {
	mock := &ParkingLotDeleter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
