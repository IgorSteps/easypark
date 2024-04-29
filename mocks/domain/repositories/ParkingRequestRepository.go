// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ParkingRequestRepository is an autogenerated mock type for the ParkingRequestRepository type
type ParkingRequestRepository struct {
	mock.Mock
}

type ParkingRequestRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ParkingRequestRepository) EXPECT() *ParkingRequestRepository_Expecter {
	return &ParkingRequestRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, parkReq
func (_m *ParkingRequestRepository) Create(ctx context.Context, parkReq *entities.ParkingRequest) error {
	ret := _m.Called(ctx, parkReq)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.ParkingRequest) error); ok {
		r0 = rf(ctx, parkReq)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ParkingRequestRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ParkingRequestRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - parkReq *entities.ParkingRequest
func (_e *ParkingRequestRepository_Expecter) Create(ctx interface{}, parkReq interface{}) *ParkingRequestRepository_Create_Call {
	return &ParkingRequestRepository_Create_Call{Call: _e.mock.On("Create", ctx, parkReq)}
}

func (_c *ParkingRequestRepository_Create_Call) Run(run func(ctx context.Context, parkReq *entities.ParkingRequest)) *ParkingRequestRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.ParkingRequest))
	})
	return _c
}

func (_c *ParkingRequestRepository_Create_Call) Return(_a0 error) *ParkingRequestRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ParkingRequestRepository_Create_Call) RunAndReturn(run func(context.Context, *entities.ParkingRequest) error) *ParkingRequestRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetMany provides a mock function with given fields: ctx, query
func (_m *ParkingRequestRepository) GetMany(ctx context.Context, query map[string]interface{}) ([]entities.ParkingRequest, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for GetMany")
	}

	var r0 []entities.ParkingRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) ([]entities.ParkingRequest, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, map[string]interface{}) []entities.ParkingRequest); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.ParkingRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, map[string]interface{}) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParkingRequestRepository_GetMany_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMany'
type ParkingRequestRepository_GetMany_Call struct {
	*mock.Call
}

// GetMany is a helper method to define mock.On call
//   - ctx context.Context
//   - query map[string]interface{}
func (_e *ParkingRequestRepository_Expecter) GetMany(ctx interface{}, query interface{}) *ParkingRequestRepository_GetMany_Call {
	return &ParkingRequestRepository_GetMany_Call{Call: _e.mock.On("GetMany", ctx, query)}
}

func (_c *ParkingRequestRepository_GetMany_Call) Run(run func(ctx context.Context, query map[string]interface{})) *ParkingRequestRepository_GetMany_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(map[string]interface{}))
	})
	return _c
}

func (_c *ParkingRequestRepository_GetMany_Call) Return(_a0 []entities.ParkingRequest, _a1 error) *ParkingRequestRepository_GetMany_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ParkingRequestRepository_GetMany_Call) RunAndReturn(run func(context.Context, map[string]interface{}) ([]entities.ParkingRequest, error)) *ParkingRequestRepository_GetMany_Call {
	_c.Call.Return(run)
	return _c
}

// GetSingle provides a mock function with given fields: ctx, id
func (_m *ParkingRequestRepository) GetSingle(ctx context.Context, id uuid.UUID) (entities.ParkingRequest, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetSingle")
	}

	var r0 entities.ParkingRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (entities.ParkingRequest, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) entities.ParkingRequest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entities.ParkingRequest)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParkingRequestRepository_GetSingle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSingle'
type ParkingRequestRepository_GetSingle_Call struct {
	*mock.Call
}

// GetSingle is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *ParkingRequestRepository_Expecter) GetSingle(ctx interface{}, id interface{}) *ParkingRequestRepository_GetSingle_Call {
	return &ParkingRequestRepository_GetSingle_Call{Call: _e.mock.On("GetSingle", ctx, id)}
}

func (_c *ParkingRequestRepository_GetSingle_Call) Run(run func(ctx context.Context, id uuid.UUID)) *ParkingRequestRepository_GetSingle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *ParkingRequestRepository_GetSingle_Call) Return(_a0 entities.ParkingRequest, _a1 error) *ParkingRequestRepository_GetSingle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ParkingRequestRepository_GetSingle_Call) RunAndReturn(run func(context.Context, uuid.UUID) (entities.ParkingRequest, error)) *ParkingRequestRepository_GetSingle_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ctx, request
func (_m *ParkingRequestRepository) Save(ctx context.Context, request *entities.ParkingRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.ParkingRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ParkingRequestRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type ParkingRequestRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - request *entities.ParkingRequest
func (_e *ParkingRequestRepository_Expecter) Save(ctx interface{}, request interface{}) *ParkingRequestRepository_Save_Call {
	return &ParkingRequestRepository_Save_Call{Call: _e.mock.On("Save", ctx, request)}
}

func (_c *ParkingRequestRepository_Save_Call) Run(run func(ctx context.Context, request *entities.ParkingRequest)) *ParkingRequestRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.ParkingRequest))
	})
	return _c
}

func (_c *ParkingRequestRepository_Save_Call) Return(_a0 error) *ParkingRequestRepository_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ParkingRequestRepository_Save_Call) RunAndReturn(run func(context.Context, *entities.ParkingRequest) error) *ParkingRequestRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// NewParkingRequestRepository creates a new instance of ParkingRequestRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewParkingRequestRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ParkingRequestRepository {
	mock := &ParkingRequestRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
