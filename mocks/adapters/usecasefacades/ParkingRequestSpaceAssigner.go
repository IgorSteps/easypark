// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ParkingRequestSpaceAssigner is an autogenerated mock type for the ParkingRequestSpaceAssigner type
type ParkingRequestSpaceAssigner struct {
	mock.Mock
}

type ParkingRequestSpaceAssigner_Expecter struct {
	mock *mock.Mock
}

func (_m *ParkingRequestSpaceAssigner) EXPECT() *ParkingRequestSpaceAssigner_Expecter {
	return &ParkingRequestSpaceAssigner_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, requestID, spaceID
func (_m *ParkingRequestSpaceAssigner) Execute(ctx context.Context, requestID uuid.UUID, spaceID uuid.UUID) error {
	ret := _m.Called(ctx, requestID, spaceID)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, requestID, spaceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ParkingRequestSpaceAssigner_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type ParkingRequestSpaceAssigner_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - requestID uuid.UUID
//   - spaceID uuid.UUID
func (_e *ParkingRequestSpaceAssigner_Expecter) Execute(ctx interface{}, requestID interface{}, spaceID interface{}) *ParkingRequestSpaceAssigner_Execute_Call {
	return &ParkingRequestSpaceAssigner_Execute_Call{Call: _e.mock.On("Execute", ctx, requestID, spaceID)}
}

func (_c *ParkingRequestSpaceAssigner_Execute_Call) Run(run func(ctx context.Context, requestID uuid.UUID, spaceID uuid.UUID)) *ParkingRequestSpaceAssigner_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(uuid.UUID))
	})
	return _c
}

func (_c *ParkingRequestSpaceAssigner_Execute_Call) Return(_a0 error) *ParkingRequestSpaceAssigner_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ParkingRequestSpaceAssigner_Execute_Call) RunAndReturn(run func(context.Context, uuid.UUID, uuid.UUID) error) *ParkingRequestSpaceAssigner_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewParkingRequestSpaceAssigner creates a new instance of ParkingRequestSpaceAssigner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewParkingRequestSpaceAssigner(t interface {
	mock.TestingT
	Cleanup(func())
}) *ParkingRequestSpaceAssigner {
	mock := &ParkingRequestSpaceAssigner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
