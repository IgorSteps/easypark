// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// NotificationFacade is an autogenerated mock type for the NotificationFacade type
type NotificationFacade struct {
	mock.Mock
}

type NotificationFacade_Expecter struct {
	mock *mock.Mock
}

func (_m *NotificationFacade) EXPECT() *NotificationFacade_Expecter {
	return &NotificationFacade_Expecter{mock: &_m.Mock}
}

// CreateNotification provides a mock function with given fields: ctx, driverID, parkingReqID, spaceID, location, notificationType
func (_m *NotificationFacade) CreateNotification(ctx context.Context, driverID uuid.UUID, parkingReqID uuid.UUID, spaceID uuid.UUID, location string, notificationType int) (entities.Notification, error) {
	ret := _m.Called(ctx, driverID, parkingReqID, spaceID, location, notificationType)

	if len(ret) == 0 {
		panic("no return value specified for CreateNotification")
	}

	var r0 entities.Notification
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, string, int) (entities.Notification, error)); ok {
		return rf(ctx, driverID, parkingReqID, spaceID, location, notificationType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, string, int) entities.Notification); ok {
		r0 = rf(ctx, driverID, parkingReqID, spaceID, location, notificationType)
	} else {
		r0 = ret.Get(0).(entities.Notification)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, string, int) error); ok {
		r1 = rf(ctx, driverID, parkingReqID, spaceID, location, notificationType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotificationFacade_CreateNotification_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateNotification'
type NotificationFacade_CreateNotification_Call struct {
	*mock.Call
}

// CreateNotification is a helper method to define mock.On call
//   - ctx context.Context
//   - driverID uuid.UUID
//   - parkingReqID uuid.UUID
//   - spaceID uuid.UUID
//   - location string
//   - notificationType int
func (_e *NotificationFacade_Expecter) CreateNotification(ctx interface{}, driverID interface{}, parkingReqID interface{}, spaceID interface{}, location interface{}, notificationType interface{}) *NotificationFacade_CreateNotification_Call {
	return &NotificationFacade_CreateNotification_Call{Call: _e.mock.On("CreateNotification", ctx, driverID, parkingReqID, spaceID, location, notificationType)}
}

func (_c *NotificationFacade_CreateNotification_Call) Run(run func(ctx context.Context, driverID uuid.UUID, parkingReqID uuid.UUID, spaceID uuid.UUID, location string, notificationType int)) *NotificationFacade_CreateNotification_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(uuid.UUID), args[3].(uuid.UUID), args[4].(string), args[5].(int))
	})
	return _c
}

func (_c *NotificationFacade_CreateNotification_Call) Return(_a0 entities.Notification, _a1 error) *NotificationFacade_CreateNotification_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NotificationFacade_CreateNotification_Call) RunAndReturn(run func(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, string, int) (entities.Notification, error)) *NotificationFacade_CreateNotification_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllNotifications provides a mock function with given fields: ctx
func (_m *NotificationFacade) GetAllNotifications(ctx context.Context) ([]entities.Notification, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllNotifications")
	}

	var r0 []entities.Notification
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entities.Notification, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entities.Notification); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entities.Notification)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NotificationFacade_GetAllNotifications_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllNotifications'
type NotificationFacade_GetAllNotifications_Call struct {
	*mock.Call
}

// GetAllNotifications is a helper method to define mock.On call
//   - ctx context.Context
func (_e *NotificationFacade_Expecter) GetAllNotifications(ctx interface{}) *NotificationFacade_GetAllNotifications_Call {
	return &NotificationFacade_GetAllNotifications_Call{Call: _e.mock.On("GetAllNotifications", ctx)}
}

func (_c *NotificationFacade_GetAllNotifications_Call) Run(run func(ctx context.Context)) *NotificationFacade_GetAllNotifications_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *NotificationFacade_GetAllNotifications_Call) Return(_a0 []entities.Notification, _a1 error) *NotificationFacade_GetAllNotifications_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NotificationFacade_GetAllNotifications_Call) RunAndReturn(run func(context.Context) ([]entities.Notification, error)) *NotificationFacade_GetAllNotifications_Call {
	_c.Call.Return(run)
	return _c
}

// NewNotificationFacade creates a new instance of NotificationFacade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNotificationFacade(t interface {
	mock.TestingT
	Cleanup(func())
}) *NotificationFacade {
	mock := &NotificationFacade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
