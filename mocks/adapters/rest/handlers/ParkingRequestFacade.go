// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// ParkingRequestFacade is an autogenerated mock type for the ParkingRequestFacade type
type ParkingRequestFacade struct {
	mock.Mock
}

type ParkingRequestFacade_Expecter struct {
	mock *mock.Mock
}

func (_m *ParkingRequestFacade) EXPECT() *ParkingRequestFacade_Expecter {
	return &ParkingRequestFacade_Expecter{mock: &_m.Mock}
}

// AssignParkingSpace provides a mock function with given fields: ctx, requestID, spaceID
func (_m *ParkingRequestFacade) AssignParkingSpace(ctx context.Context, requestID uuid.UUID, spaceID uuid.UUID) error {
	ret := _m.Called(ctx, requestID, spaceID)

	if len(ret) == 0 {
		panic("no return value specified for AssignParkingSpace")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, requestID, spaceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ParkingRequestFacade_AssignParkingSpace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AssignParkingSpace'
type ParkingRequestFacade_AssignParkingSpace_Call struct {
	*mock.Call
}

// AssignParkingSpace is a helper method to define mock.On call
//   - ctx context.Context
//   - requestID uuid.UUID
//   - spaceID uuid.UUID
func (_e *ParkingRequestFacade_Expecter) AssignParkingSpace(ctx interface{}, requestID interface{}, spaceID interface{}) *ParkingRequestFacade_AssignParkingSpace_Call {
	return &ParkingRequestFacade_AssignParkingSpace_Call{Call: _e.mock.On("AssignParkingSpace", ctx, requestID, spaceID)}
}

func (_c *ParkingRequestFacade_AssignParkingSpace_Call) Run(run func(ctx context.Context, requestID uuid.UUID, spaceID uuid.UUID)) *ParkingRequestFacade_AssignParkingSpace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(uuid.UUID))
	})
	return _c
}

func (_c *ParkingRequestFacade_AssignParkingSpace_Call) Return(_a0 error) *ParkingRequestFacade_AssignParkingSpace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ParkingRequestFacade_AssignParkingSpace_Call) RunAndReturn(run func(context.Context, uuid.UUID, uuid.UUID) error) *ParkingRequestFacade_AssignParkingSpace_Call {
	_c.Call.Return(run)
	return _c
}

// CreateParkingRequest provides a mock function with given fields: ctx, parkingRequest
func (_m *ParkingRequestFacade) CreateParkingRequest(ctx context.Context, parkingRequest *entities.ParkingRequest) (*entities.ParkingRequest, error) {
	ret := _m.Called(ctx, parkingRequest)

	if len(ret) == 0 {
		panic("no return value specified for CreateParkingRequest")
	}

	var r0 *entities.ParkingRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entities.ParkingRequest) (*entities.ParkingRequest, error)); ok {
		return rf(ctx, parkingRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entities.ParkingRequest) *entities.ParkingRequest); ok {
		r0 = rf(ctx, parkingRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.ParkingRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entities.ParkingRequest) error); ok {
		r1 = rf(ctx, parkingRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParkingRequestFacade_CreateParkingRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateParkingRequest'
type ParkingRequestFacade_CreateParkingRequest_Call struct {
	*mock.Call
}

// CreateParkingRequest is a helper method to define mock.On call
//   - ctx context.Context
//   - parkingRequest *entities.ParkingRequest
func (_e *ParkingRequestFacade_Expecter) CreateParkingRequest(ctx interface{}, parkingRequest interface{}) *ParkingRequestFacade_CreateParkingRequest_Call {
	return &ParkingRequestFacade_CreateParkingRequest_Call{Call: _e.mock.On("CreateParkingRequest", ctx, parkingRequest)}
}

func (_c *ParkingRequestFacade_CreateParkingRequest_Call) Run(run func(ctx context.Context, parkingRequest *entities.ParkingRequest)) *ParkingRequestFacade_CreateParkingRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*entities.ParkingRequest))
	})
	return _c
}

func (_c *ParkingRequestFacade_CreateParkingRequest_Call) Return(_a0 *entities.ParkingRequest, _a1 error) *ParkingRequestFacade_CreateParkingRequest_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ParkingRequestFacade_CreateParkingRequest_Call) RunAndReturn(run func(context.Context, *entities.ParkingRequest) (*entities.ParkingRequest, error)) *ParkingRequestFacade_CreateParkingRequest_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllParkingRequests provides a mock function with given fields: ctx
func (_m *ParkingRequestFacade) GetAllParkingRequests(ctx context.Context) ([]entities.ParkingRequest, error) {
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

// ParkingRequestFacade_GetAllParkingRequests_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllParkingRequests'
type ParkingRequestFacade_GetAllParkingRequests_Call struct {
	*mock.Call
}

// GetAllParkingRequests is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ParkingRequestFacade_Expecter) GetAllParkingRequests(ctx interface{}) *ParkingRequestFacade_GetAllParkingRequests_Call {
	return &ParkingRequestFacade_GetAllParkingRequests_Call{Call: _e.mock.On("GetAllParkingRequests", ctx)}
}

func (_c *ParkingRequestFacade_GetAllParkingRequests_Call) Run(run func(ctx context.Context)) *ParkingRequestFacade_GetAllParkingRequests_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ParkingRequestFacade_GetAllParkingRequests_Call) Return(_a0 []entities.ParkingRequest, _a1 error) *ParkingRequestFacade_GetAllParkingRequests_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ParkingRequestFacade_GetAllParkingRequests_Call) RunAndReturn(run func(context.Context) ([]entities.ParkingRequest, error)) *ParkingRequestFacade_GetAllParkingRequests_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateParkingRequestStatus provides a mock function with given fields: ctx, id, status
func (_m *ParkingRequestFacade) UpdateParkingRequestStatus(ctx context.Context, id uuid.UUID, status string) error {
	ret := _m.Called(ctx, id, status)

	if len(ret) == 0 {
		panic("no return value specified for UpdateParkingRequestStatus")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) error); ok {
		r0 = rf(ctx, id, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ParkingRequestFacade_UpdateParkingRequestStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateParkingRequestStatus'
type ParkingRequestFacade_UpdateParkingRequestStatus_Call struct {
	*mock.Call
}

// UpdateParkingRequestStatus is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
//   - status string
func (_e *ParkingRequestFacade_Expecter) UpdateParkingRequestStatus(ctx interface{}, id interface{}, status interface{}) *ParkingRequestFacade_UpdateParkingRequestStatus_Call {
	return &ParkingRequestFacade_UpdateParkingRequestStatus_Call{Call: _e.mock.On("UpdateParkingRequestStatus", ctx, id, status)}
}

func (_c *ParkingRequestFacade_UpdateParkingRequestStatus_Call) Run(run func(ctx context.Context, id uuid.UUID, status string)) *ParkingRequestFacade_UpdateParkingRequestStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(string))
	})
	return _c
}

func (_c *ParkingRequestFacade_UpdateParkingRequestStatus_Call) Return(_a0 error) *ParkingRequestFacade_UpdateParkingRequestStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ParkingRequestFacade_UpdateParkingRequestStatus_Call) RunAndReturn(run func(context.Context, uuid.UUID, string) error) *ParkingRequestFacade_UpdateParkingRequestStatus_Call {
	_c.Call.Return(run)
	return _c
}

// NewParkingRequestFacade creates a new instance of ParkingRequestFacade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewParkingRequestFacade(t interface {
	mock.TestingT
	Cleanup(func())
}) *ParkingRequestFacade {
	mock := &ParkingRequestFacade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
