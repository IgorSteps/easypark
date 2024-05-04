// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"
	mock "github.com/stretchr/testify/mock"
)

// ParkingLotCreator is an autogenerated mock type for the ParkingLotCreator type
type ParkingLotCreator struct {
	mock.Mock
}

type ParkingLotCreator_Expecter struct {
	mock *mock.Mock
}

func (_m *ParkingLotCreator) EXPECT() *ParkingLotCreator_Expecter {
	return &ParkingLotCreator_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, name, capacity
func (_m *ParkingLotCreator) Execute(ctx context.Context, name string, capacity int) (entities.ParkingLot, error) {
	ret := _m.Called(ctx, name, capacity)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 entities.ParkingLot
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) (entities.ParkingLot, error)); ok {
		return rf(ctx, name, capacity)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int) entities.ParkingLot); ok {
		r0 = rf(ctx, name, capacity)
	} else {
		r0 = ret.Get(0).(entities.ParkingLot)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, name, capacity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParkingLotCreator_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type ParkingLotCreator_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
//   - capacity int
func (_e *ParkingLotCreator_Expecter) Execute(ctx interface{}, name interface{}, capacity interface{}) *ParkingLotCreator_Execute_Call {
	return &ParkingLotCreator_Execute_Call{Call: _e.mock.On("Execute", ctx, name, capacity)}
}

func (_c *ParkingLotCreator_Execute_Call) Run(run func(ctx context.Context, name string, capacity int)) *ParkingLotCreator_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(int))
	})
	return _c
}

func (_c *ParkingLotCreator_Execute_Call) Return(_a0 entities.ParkingLot, _a1 error) *ParkingLotCreator_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ParkingLotCreator_Execute_Call) RunAndReturn(run func(context.Context, string, int) (entities.ParkingLot, error)) *ParkingLotCreator_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewParkingLotCreator creates a new instance of ParkingLotCreator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewParkingLotCreator(t interface {
	mock.TestingT
	Cleanup(func())
}) *ParkingLotCreator {
	mock := &ParkingLotCreator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
