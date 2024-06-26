// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ParkingLotSingleGetter is an autogenerated mock type for the ParkingLotSingleGetter type
type ParkingLotSingleGetter struct {
	mock.Mock
}

type ParkingLotSingleGetter_Expecter struct {
	mock *mock.Mock
}

func (_m *ParkingLotSingleGetter) EXPECT() *ParkingLotSingleGetter_Expecter {
	return &ParkingLotSingleGetter_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, id
func (_m *ParkingLotSingleGetter) Execute(ctx context.Context, id uuid.UUID) (*entities.ParkingLot, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 *entities.ParkingLot
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*entities.ParkingLot, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entities.ParkingLot); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ParkingLot)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParkingLotSingleGetter_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type ParkingLotSingleGetter_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *ParkingLotSingleGetter_Expecter) Execute(ctx interface{}, id interface{}) *ParkingLotSingleGetter_Execute_Call {
	return &ParkingLotSingleGetter_Execute_Call{Call: _e.mock.On("Execute", ctx, id)}
}

func (_c *ParkingLotSingleGetter_Execute_Call) Run(run func(ctx context.Context, id uuid.UUID)) *ParkingLotSingleGetter_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *ParkingLotSingleGetter_Execute_Call) Return(_a0 *entities.ParkingLot, _a1 error) *ParkingLotSingleGetter_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ParkingLotSingleGetter_Execute_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*entities.ParkingLot, error)) *ParkingLotSingleGetter_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewParkingLotSingleGetter creates a new instance of ParkingLotSingleGetter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewParkingLotSingleGetter(t interface {
	mock.TestingT
	Cleanup(func())
}) *ParkingLotSingleGetter {
	mock := &ParkingLotSingleGetter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
