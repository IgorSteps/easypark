// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// MessageFacade is an autogenerated mock type for the MessageFacade type
type MessageFacade struct {
	mock.Mock
}

type MessageFacade_Expecter struct {
	mock *mock.Mock
}

func (_m *MessageFacade) EXPECT() *MessageFacade_Expecter {
	return &MessageFacade_Expecter{mock: &_m.Mock}
}

// CreateMessage provides a mock function with given fields: ctx, msg
func (_m *MessageFacade) CreateMessage(ctx context.Context, msg entities.Message) (entities.Message, error) {
	ret := _m.Called(ctx, msg)

	if len(ret) == 0 {
		panic("no return value specified for CreateMessage")
	}

	var r0 entities.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entities.Message) (entities.Message, error)); ok {
		return rf(ctx, msg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entities.Message) entities.Message); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Get(0).(entities.Message)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entities.Message) error); ok {
		r1 = rf(ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MessageFacade_CreateMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMessage'
type MessageFacade_CreateMessage_Call struct {
	*mock.Call
}

// CreateMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - msg entities.Message
func (_e *MessageFacade_Expecter) CreateMessage(ctx interface{}, msg interface{}) *MessageFacade_CreateMessage_Call {
	return &MessageFacade_CreateMessage_Call{Call: _e.mock.On("CreateMessage", ctx, msg)}
}

func (_c *MessageFacade_CreateMessage_Call) Run(run func(ctx context.Context, msg entities.Message)) *MessageFacade_CreateMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(entities.Message))
	})
	return _c
}

func (_c *MessageFacade_CreateMessage_Call) Return(_a0 entities.Message, _a1 error) *MessageFacade_CreateMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MessageFacade_CreateMessage_Call) RunAndReturn(run func(context.Context, entities.Message) (entities.Message, error)) *MessageFacade_CreateMessage_Call {
	_c.Call.Return(run)
	return _c
}

// NewMessageFacade creates a new instance of MessageFacade. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMessageFacade(t interface {
	mock.TestingT
	Cleanup(func())
}) *MessageFacade {
	mock := &MessageFacade{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
