// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/IgorSteps/easypark/internal/domain/entities"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MessageEnqueuer is an autogenerated mock type for the MessageEnqueuer type
type MessageEnqueuer struct {
	mock.Mock
}

type MessageEnqueuer_Expecter struct {
	mock *mock.Mock
}

func (_m *MessageEnqueuer) EXPECT() *MessageEnqueuer_Expecter {
	return &MessageEnqueuer_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, senderID, receiverID, content
func (_m *MessageEnqueuer) Execute(ctx context.Context, senderID uuid.UUID, receiverID uuid.UUID, content string) (entities.Message, error) {
	ret := _m.Called(ctx, senderID, receiverID, content)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 entities.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID, string) (entities.Message, error)); ok {
		return rf(ctx, senderID, receiverID, content)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID, string) entities.Message); ok {
		r0 = rf(ctx, senderID, receiverID, content)
	} else {
		r0 = ret.Get(0).(entities.Message)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, uuid.UUID, string) error); ok {
		r1 = rf(ctx, senderID, receiverID, content)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MessageEnqueuer_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MessageEnqueuer_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - senderID uuid.UUID
//   - receiverID uuid.UUID
//   - content string
func (_e *MessageEnqueuer_Expecter) Execute(ctx interface{}, senderID interface{}, receiverID interface{}, content interface{}) *MessageEnqueuer_Execute_Call {
	return &MessageEnqueuer_Execute_Call{Call: _e.mock.On("Execute", ctx, senderID, receiverID, content)}
}

func (_c *MessageEnqueuer_Execute_Call) Run(run func(ctx context.Context, senderID uuid.UUID, receiverID uuid.UUID, content string)) *MessageEnqueuer_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(uuid.UUID), args[3].(string))
	})
	return _c
}

func (_c *MessageEnqueuer_Execute_Call) Return(_a0 entities.Message, _a1 error) *MessageEnqueuer_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MessageEnqueuer_Execute_Call) RunAndReturn(run func(context.Context, uuid.UUID, uuid.UUID, string) (entities.Message, error)) *MessageEnqueuer_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMessageEnqueuer creates a new instance of MessageEnqueuer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMessageEnqueuer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MessageEnqueuer {
	mock := &MessageEnqueuer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
