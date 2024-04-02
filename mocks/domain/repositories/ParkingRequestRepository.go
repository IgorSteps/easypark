// Code generated by mockery v2.42.0. DO NOT EDIT.

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

// CreateParkingRequest provides a mock function with given fields: ctx, parkReq
func (_m *ParkingRequestRepository) CreateParkingRequest(ctx context.Context, parkReq *entities.ParkingRequest) error {
	ret := _m.Called(ctx, parkReq)

	if len(ret) == 0 {
		panic("no return value specified for CreateParkingRequest")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.ParkingRequest) error); ok {
		r0 = rf(ctx, parkReq)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ParkingRequestRepository_CreateParkingRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateParkingRequest'
type ParkingRequestRepository_CreateParkingRequest_Call struct {
	*mock.Call
}

// CreateParkingRequest is a helper method to define mock.On call
//   - ctx context.Context
//   - parkReq *entities.ParkingRequest
func (_e *ParkingRequestRepository_Expecter) CreateParkingRequest(ctx interface{}, parkReq interface{}) *ParkingRequestRepository_CreateParkingRequest_Call {
	return &ParkingRequestRepository_CreateParkingRequest_Call{Call: _e.mock.On("CreateParkingRequest", ctx, parkReq)}
}

func (_c *ParkingRequestRepository_CreateParkingRequest_Call) Run(run func(ctx context.Context, parkReq *entities.ParkingRequest)) *ParkingRequestRepository_CreateParkingRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.ParkingRequest))
	})
	return _c
}

func (_c *ParkingRequestRepository_CreateParkingRequest_Call) Return(_a0 error) *ParkingRequestRepository_CreateParkingRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ParkingRequestRepository_CreateParkingRequest_Call) RunAndReturn(run func(context.Context, *entities.ParkingRequest) error) *ParkingRequestRepository_CreateParkingRequest_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllParkingRequests provides a mock function with given fields: ctx
func (_m *ParkingRequestRepository) GetAllParkingRequests(ctx context.Context) ([]entities.ParkingRequest, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllParkingRequests")
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

// ParkingRequestRepository_GetAllParkingRequests_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllParkingRequests'
type ParkingRequestRepository_GetAllParkingRequests_Call struct {
	*mock.Call
}

// GetAllParkingRequests is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ParkingRequestRepository_Expecter) GetAllParkingRequests(ctx interface{}) *ParkingRequestRepository_GetAllParkingRequests_Call {
	return &ParkingRequestRepository_GetAllParkingRequests_Call{Call: _e.mock.On("GetAllParkingRequests", ctx)}
}

func (_c *ParkingRequestRepository_GetAllParkingRequests_Call) Run(run func(ctx context.Context)) *ParkingRequestRepository_GetAllParkingRequests_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ParkingRequestRepository_GetAllParkingRequests_Call) Return(_a0 []entities.ParkingRequest, _a1 error) *ParkingRequestRepository_GetAllParkingRequests_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ParkingRequestRepository_GetAllParkingRequests_Call) RunAndReturn(run func(context.Context) ([]entities.ParkingRequest, error)) *ParkingRequestRepository_GetAllParkingRequests_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllParkingRequestsForUser provides a mock function with given fields: ctx, userID
func (_m *ParkingRequestRepository) GetAllParkingRequestsForUser(ctx context.Context, userID uuid.UUID) ([]entities.ParkingRequest, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetAllParkingRequestsForUser")
	}

	var r0 []entities.ParkingRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]entities.ParkingRequest, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []entities.ParkingRequest); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.ParkingRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParkingRequestRepository_GetAllParkingRequestsForUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllParkingRequestsForUser'
type ParkingRequestRepository_GetAllParkingRequestsForUser_Call struct {
	*mock.Call
}

// GetAllParkingRequestsForUser is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uuid.UUID
func (_e *ParkingRequestRepository_Expecter) GetAllParkingRequestsForUser(ctx interface{}, userID interface{}) *ParkingRequestRepository_GetAllParkingRequestsForUser_Call {
	return &ParkingRequestRepository_GetAllParkingRequestsForUser_Call{Call: _e.mock.On("GetAllParkingRequestsForUser", ctx, userID)}
}

func (_c *ParkingRequestRepository_GetAllParkingRequestsForUser_Call) Run(run func(ctx context.Context, userID uuid.UUID)) *ParkingRequestRepository_GetAllParkingRequestsForUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *ParkingRequestRepository_GetAllParkingRequestsForUser_Call) Return(_a0 []entities.ParkingRequest, _a1 error) *ParkingRequestRepository_GetAllParkingRequestsForUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ParkingRequestRepository_GetAllParkingRequestsForUser_Call) RunAndReturn(run func(context.Context, uuid.UUID) ([]entities.ParkingRequest, error)) *ParkingRequestRepository_GetAllParkingRequestsForUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetParkingRequestByID provides a mock function with given fields: ctx, id
func (_m *ParkingRequestRepository) GetParkingRequestByID(ctx context.Context, id uuid.UUID) (entities.ParkingRequest, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetParkingRequestByID")
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

// ParkingRequestRepository_GetParkingRequestByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetParkingRequestByID'
type ParkingRequestRepository_GetParkingRequestByID_Call struct {
	*mock.Call
}

// GetParkingRequestByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *ParkingRequestRepository_Expecter) GetParkingRequestByID(ctx interface{}, id interface{}) *ParkingRequestRepository_GetParkingRequestByID_Call {
	return &ParkingRequestRepository_GetParkingRequestByID_Call{Call: _e.mock.On("GetParkingRequestByID", ctx, id)}
}

func (_c *ParkingRequestRepository_GetParkingRequestByID_Call) Run(run func(ctx context.Context, id uuid.UUID)) *ParkingRequestRepository_GetParkingRequestByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *ParkingRequestRepository_GetParkingRequestByID_Call) Return(_a0 entities.ParkingRequest, _a1 error) *ParkingRequestRepository_GetParkingRequestByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ParkingRequestRepository_GetParkingRequestByID_Call) RunAndReturn(run func(context.Context, uuid.UUID) (entities.ParkingRequest, error)) *ParkingRequestRepository_GetParkingRequestByID_Call {
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
