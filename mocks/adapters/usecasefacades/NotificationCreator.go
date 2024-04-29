// Code generated by mockery v2.42.3. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// NotificationCreator is an autogenerated mock type for the NotificationCreator type
type NotificationCreator struct {
	mock.Mock
}

type NotificationCreator_Expecter struct {
	mock *mock.Mock
}

func (_m *NotificationCreator) EXPECT() *NotificationCreator_Expecter {
	return &NotificationCreator_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, driverID, parkingReqID, spaceID, location, notificationType
func (_m *NotificationCreator) Execute(ctx context.Context, driverID uuid.UUID, parkingReqID uuid.UUID, spaceID uuid.UUID, location string, notificationType int) (entities.Notification, error) {
	ret := _m.Called(ctx, driverID, parkingReqID, spaceID, location, notificationType)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
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

// NotificationCreator_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type NotificationCreator_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - driverID uuid.UUID
//   - parkingReqID uuid.UUID
//   - spaceID uuid.UUID
//   - location string
//   - notificationType int
func (_e *NotificationCreator_Expecter) Execute(ctx interface{}, driverID interface{}, parkingReqID interface{}, spaceID interface{}, location interface{}, notificationType interface{}) *NotificationCreator_Execute_Call {
	return &NotificationCreator_Execute_Call{Call: _e.mock.On("Execute", ctx, driverID, parkingReqID, spaceID, location, notificationType)}
}

func (_c *NotificationCreator_Execute_Call) Run(run func(ctx context.Context, driverID uuid.UUID, parkingReqID uuid.UUID, spaceID uuid.UUID, location string, notificationType int)) *NotificationCreator_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(uuid.UUID), args[3].(uuid.UUID), args[4].(string), args[5].(int))
	})
	return _c
}

func (_c *NotificationCreator_Execute_Call) Return(_a0 entities.Notification, _a1 error) *NotificationCreator_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *NotificationCreator_Execute_Call) RunAndReturn(run func(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, string, int) (entities.Notification, error)) *NotificationCreator_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewNotificationCreator creates a new instance of NotificationCreator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNotificationCreator(t interface {
	mock.TestingT
	Cleanup(func())
}) *NotificationCreator {
	mock := &NotificationCreator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
