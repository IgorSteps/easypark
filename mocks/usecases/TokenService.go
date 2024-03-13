// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	entities "github.com/IgorSteps/easypark/internal/domain/entities"
	mock "github.com/stretchr/testify/mock"
)

// TokenService is an autogenerated mock type for the TokenService type
type TokenService struct {
	mock.Mock
}

type TokenService_Expecter struct {
	mock *mock.Mock
}

func (_m *TokenService) EXPECT() *TokenService_Expecter {
	return &TokenService_Expecter{mock: &_m.Mock}
}

// GenerateToken provides a mock function with given fields: user, expiresAt
func (_m *TokenService) GenerateToken(user *entities.User, expiresAt int64) (string, error) {
	ret := _m.Called(user, expiresAt)

	if len(ret) == 0 {
		panic("no return value specified for GenerateToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*entities.User, int64) (string, error)); ok {
		return rf(user, expiresAt)
	}
	if rf, ok := ret.Get(0).(func(*entities.User, int64) string); ok {
		r0 = rf(user, expiresAt)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*entities.User, int64) error); ok {
		r1 = rf(user, expiresAt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TokenService_GenerateToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GenerateToken'
type TokenService_GenerateToken_Call struct {
	*mock.Call
}

// GenerateToken is a helper method to define mock.On call
//   - user *entities.User
//   - expiresAt int64
func (_e *TokenService_Expecter) GenerateToken(user interface{}, expiresAt interface{}) *TokenService_GenerateToken_Call {
	return &TokenService_GenerateToken_Call{Call: _e.mock.On("GenerateToken", user, expiresAt)}
}

func (_c *TokenService_GenerateToken_Call) Run(run func(user *entities.User, expiresAt int64)) *TokenService_GenerateToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entities.User), args[1].(int64))
	})
	return _c
}

func (_c *TokenService_GenerateToken_Call) Return(_a0 string, _a1 error) *TokenService_GenerateToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TokenService_GenerateToken_Call) RunAndReturn(run func(*entities.User, int64) (string, error)) *TokenService_GenerateToken_Call {
	_c.Call.Return(run)
	return _c
}

// NewTokenService creates a new instance of TokenService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTokenService(t interface {
	mock.TestingT
	Cleanup(func())
}) *TokenService {
	mock := &TokenService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}