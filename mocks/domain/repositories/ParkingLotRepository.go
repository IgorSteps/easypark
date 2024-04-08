// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"
	mock "github.com/stretchr/testify/mock"
)

// ParkingLotRepository is an autogenerated mock type for the ParkingLotRepository type
type ParkingLotRepository struct {
	mock.Mock
}

type ParkingLotRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ParkingLotRepository) EXPECT() *ParkingLotRepository_Expecter {
	return &ParkingLotRepository_Expecter{mock: &_m.Mock}
}

// CreateParkingLot provides a mock function with given fields: ctx, parkingLot
func (_m *ParkingLotRepository) CreateParkingLot(ctx context.Context, parkingLot *entities.ParkingLot) error {
	ret := _m.Called(ctx, parkingLot)

	if len(ret) == 0 {
		panic("no return value specified for CreateParkingLot")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.ParkingLot) error); ok {
		r0 = rf(ctx, parkingLot)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ParkingLotRepository_CreateParkingLot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateParkingLot'
type ParkingLotRepository_CreateParkingLot_Call struct {
	*mock.Call
}

// CreateParkingLot is a helper method to define mock.On call
//   - ctx context.Context
//   - parkingLot *entities.ParkingLot
func (_e *ParkingLotRepository_Expecter) CreateParkingLot(ctx interface{}, parkingLot interface{}) *ParkingLotRepository_CreateParkingLot_Call {
	return &ParkingLotRepository_CreateParkingLot_Call{Call: _e.mock.On("CreateParkingLot", ctx, parkingLot)}
}

func (_c *ParkingLotRepository_CreateParkingLot_Call) Run(run func(ctx context.Context, parkingLot *entities.ParkingLot)) *ParkingLotRepository_CreateParkingLot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.ParkingLot))
	})
	return _c
}

func (_c *ParkingLotRepository_CreateParkingLot_Call) Return(_a0 error) *ParkingLotRepository_CreateParkingLot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ParkingLotRepository_CreateParkingLot_Call) RunAndReturn(run func(context.Context, *entities.ParkingLot) error) *ParkingLotRepository_CreateParkingLot_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllParkingLots provides a mock function with given fields: ctx
func (_m *ParkingLotRepository) GetAllParkingLots(ctx context.Context) ([]entities.ParkingLot, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllParkingLots")
	}

	var r0 []entities.ParkingLot
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entities.ParkingLot, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entities.ParkingLot); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.ParkingLot)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParkingLotRepository_GetAllParkingLots_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllParkingLots'
type ParkingLotRepository_GetAllParkingLots_Call struct {
	*mock.Call
}

// GetAllParkingLots is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ParkingLotRepository_Expecter) GetAllParkingLots(ctx interface{}) *ParkingLotRepository_GetAllParkingLots_Call {
	return &ParkingLotRepository_GetAllParkingLots_Call{Call: _e.mock.On("GetAllParkingLots", ctx)}
}

func (_c *ParkingLotRepository_GetAllParkingLots_Call) Run(run func(ctx context.Context)) *ParkingLotRepository_GetAllParkingLots_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ParkingLotRepository_GetAllParkingLots_Call) Return(_a0 []entities.ParkingLot, _a1 error) *ParkingLotRepository_GetAllParkingLots_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ParkingLotRepository_GetAllParkingLots_Call) RunAndReturn(run func(context.Context) ([]entities.ParkingLot, error)) *ParkingLotRepository_GetAllParkingLots_Call {
	_c.Call.Return(run)
	return _c
}

// NewParkingLotRepository creates a new instance of ParkingLotRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewParkingLotRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ParkingLotRepository {
	mock := &ParkingLotRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
