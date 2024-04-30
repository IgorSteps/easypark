// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ParkingRequestStatusUpdater is an autogenerated mock type for the ParkingRequestStatusUpdater type
type ParkingRequestStatusUpdater struct {
	mock.Mock
}

type ParkingRequestStatusUpdater_Expecter struct {
	mock *mock.Mock
}

func (_m *ParkingRequestStatusUpdater) EXPECT() *ParkingRequestStatusUpdater_Expecter {
	return &ParkingRequestStatusUpdater_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, id, status
func (_m *ParkingRequestStatusUpdater) Execute(ctx context.Context, id uuid.UUID, status string) error {
	ret := _m.Called(ctx, id, status)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) error); ok {
		r0 = rf(ctx, id, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ParkingRequestStatusUpdater_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type ParkingRequestStatusUpdater_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
//   - status string
func (_e *ParkingRequestStatusUpdater_Expecter) Execute(ctx interface{}, id interface{}, status interface{}) *ParkingRequestStatusUpdater_Execute_Call {
	return &ParkingRequestStatusUpdater_Execute_Call{Call: _e.mock.On("Execute", ctx, id, status)}
}

func (_c *ParkingRequestStatusUpdater_Execute_Call) Run(run func(ctx context.Context, id uuid.UUID, status string)) *ParkingRequestStatusUpdater_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(string))
	})
	return _c
}

func (_c *ParkingRequestStatusUpdater_Execute_Call) Return(_a0 error) *ParkingRequestStatusUpdater_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ParkingRequestStatusUpdater_Execute_Call) RunAndReturn(run func(context.Context, uuid.UUID, string) error) *ParkingRequestStatusUpdater_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewParkingRequestStatusUpdater creates a new instance of ParkingRequestStatusUpdater. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewParkingRequestStatusUpdater(t interface {
	mock.TestingT
	Cleanup(func())
}) *ParkingRequestStatusUpdater {
	mock := &ParkingRequestStatusUpdater{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
