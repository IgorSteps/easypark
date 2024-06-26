// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// HandlerFactory is an autogenerated mock type for the HandlerFactory type
type HandlerFactory struct {
	mock.Mock
}

type HandlerFactory_Expecter struct {
	mock *mock.Mock
}

func (_m *HandlerFactory) EXPECT() *HandlerFactory_Expecter {
	return &HandlerFactory_Expecter{mock: &_m.Mock}
}

// AssignParkingSpace provides a mock function with given fields:
func (_m *HandlerFactory) AssignParkingSpace() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AssignParkingSpace")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_AssignParkingSpace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AssignParkingSpace'
type HandlerFactory_AssignParkingSpace_Call struct {
	*mock.Call
}

// AssignParkingSpace is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) AssignParkingSpace() *HandlerFactory_AssignParkingSpace_Call {
	return &HandlerFactory_AssignParkingSpace_Call{Call: _e.mock.On("AssignParkingSpace")}
}

func (_c *HandlerFactory_AssignParkingSpace_Call) Run(run func()) *HandlerFactory_AssignParkingSpace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_AssignParkingSpace_Call) Return(_a0 http.Handler) *HandlerFactory_AssignParkingSpace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_AssignParkingSpace_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_AssignParkingSpace_Call {
	_c.Call.Return(run)
	return _c
}

// AutomaticallyAssignParkingSpace provides a mock function with given fields:
func (_m *HandlerFactory) AutomaticallyAssignParkingSpace() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AutomaticallyAssignParkingSpace")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_AutomaticallyAssignParkingSpace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AutomaticallyAssignParkingSpace'
type HandlerFactory_AutomaticallyAssignParkingSpace_Call struct {
	*mock.Call
}

// AutomaticallyAssignParkingSpace is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) AutomaticallyAssignParkingSpace() *HandlerFactory_AutomaticallyAssignParkingSpace_Call {
	return &HandlerFactory_AutomaticallyAssignParkingSpace_Call{Call: _e.mock.On("AutomaticallyAssignParkingSpace")}
}

func (_c *HandlerFactory_AutomaticallyAssignParkingSpace_Call) Run(run func()) *HandlerFactory_AutomaticallyAssignParkingSpace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_AutomaticallyAssignParkingSpace_Call) Return(_a0 http.Handler) *HandlerFactory_AutomaticallyAssignParkingSpace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_AutomaticallyAssignParkingSpace_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_AutomaticallyAssignParkingSpace_Call {
	_c.Call.Return(run)
	return _c
}

// CheckForLateArrivals provides a mock function with given fields:
func (_m *HandlerFactory) CheckForLateArrivals() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CheckForLateArrivals")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_CheckForLateArrivals_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckForLateArrivals'
type HandlerFactory_CheckForLateArrivals_Call struct {
	*mock.Call
}

// CheckForLateArrivals is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) CheckForLateArrivals() *HandlerFactory_CheckForLateArrivals_Call {
	return &HandlerFactory_CheckForLateArrivals_Call{Call: _e.mock.On("CheckForLateArrivals")}
}

func (_c *HandlerFactory_CheckForLateArrivals_Call) Run(run func()) *HandlerFactory_CheckForLateArrivals_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_CheckForLateArrivals_Call) Return(_a0 http.Handler) *HandlerFactory_CheckForLateArrivals_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_CheckForLateArrivals_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_CheckForLateArrivals_Call {
	_c.Call.Return(run)
	return _c
}

// CheckForOverStays provides a mock function with given fields:
func (_m *HandlerFactory) CheckForOverStays() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CheckForOverStays")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_CheckForOverStays_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckForOverStays'
type HandlerFactory_CheckForOverStays_Call struct {
	*mock.Call
}

// CheckForOverStays is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) CheckForOverStays() *HandlerFactory_CheckForOverStays_Call {
	return &HandlerFactory_CheckForOverStays_Call{Call: _e.mock.On("CheckForOverStays")}
}

func (_c *HandlerFactory_CheckForOverStays_Call) Run(run func()) *HandlerFactory_CheckForOverStays_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_CheckForOverStays_Call) Return(_a0 http.Handler) *HandlerFactory_CheckForOverStays_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_CheckForOverStays_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_CheckForOverStays_Call {
	_c.Call.Return(run)
	return _c
}

// CreateNotification provides a mock function with given fields:
func (_m *HandlerFactory) CreateNotification() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CreateNotification")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_CreateNotification_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateNotification'
type HandlerFactory_CreateNotification_Call struct {
	*mock.Call
}

// CreateNotification is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) CreateNotification() *HandlerFactory_CreateNotification_Call {
	return &HandlerFactory_CreateNotification_Call{Call: _e.mock.On("CreateNotification")}
}

func (_c *HandlerFactory_CreateNotification_Call) Run(run func()) *HandlerFactory_CreateNotification_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_CreateNotification_Call) Return(_a0 http.Handler) *HandlerFactory_CreateNotification_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_CreateNotification_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_CreateNotification_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteParkingLot provides a mock function with given fields:
func (_m *HandlerFactory) DeleteParkingLot() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DeleteParkingLot")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_DeleteParkingLot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteParkingLot'
type HandlerFactory_DeleteParkingLot_Call struct {
	*mock.Call
}

// DeleteParkingLot is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) DeleteParkingLot() *HandlerFactory_DeleteParkingLot_Call {
	return &HandlerFactory_DeleteParkingLot_Call{Call: _e.mock.On("DeleteParkingLot")}
}

func (_c *HandlerFactory_DeleteParkingLot_Call) Run(run func()) *HandlerFactory_DeleteParkingLot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_DeleteParkingLot_Call) Return(_a0 http.Handler) *HandlerFactory_DeleteParkingLot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_DeleteParkingLot_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_DeleteParkingLot_Call {
	_c.Call.Return(run)
	return _c
}

// DriverBan provides a mock function with given fields:
func (_m *HandlerFactory) DriverBan() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DriverBan")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_DriverBan_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DriverBan'
type HandlerFactory_DriverBan_Call struct {
	*mock.Call
}

// DriverBan is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) DriverBan() *HandlerFactory_DriverBan_Call {
	return &HandlerFactory_DriverBan_Call{Call: _e.mock.On("DriverBan")}
}

func (_c *HandlerFactory_DriverBan_Call) Run(run func()) *HandlerFactory_DriverBan_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_DriverBan_Call) Return(_a0 http.Handler) *HandlerFactory_DriverBan_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_DriverBan_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_DriverBan_Call {
	_c.Call.Return(run)
	return _c
}

// DriverCreate provides a mock function with given fields:
func (_m *HandlerFactory) DriverCreate() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for DriverCreate")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_DriverCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DriverCreate'
type HandlerFactory_DriverCreate_Call struct {
	*mock.Call
}

// DriverCreate is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) DriverCreate() *HandlerFactory_DriverCreate_Call {
	return &HandlerFactory_DriverCreate_Call{Call: _e.mock.On("DriverCreate")}
}

func (_c *HandlerFactory_DriverCreate_Call) Run(run func()) *HandlerFactory_DriverCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_DriverCreate_Call) Return(_a0 http.Handler) *HandlerFactory_DriverCreate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_DriverCreate_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_DriverCreate_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllAlerts provides a mock function with given fields:
func (_m *HandlerFactory) GetAllAlerts() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllAlerts")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_GetAllAlerts_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllAlerts'
type HandlerFactory_GetAllAlerts_Call struct {
	*mock.Call
}

// GetAllAlerts is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) GetAllAlerts() *HandlerFactory_GetAllAlerts_Call {
	return &HandlerFactory_GetAllAlerts_Call{Call: _e.mock.On("GetAllAlerts")}
}

func (_c *HandlerFactory_GetAllAlerts_Call) Run(run func()) *HandlerFactory_GetAllAlerts_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_GetAllAlerts_Call) Return(_a0 http.Handler) *HandlerFactory_GetAllAlerts_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_GetAllAlerts_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_GetAllAlerts_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllDrivers provides a mock function with given fields:
func (_m *HandlerFactory) GetAllDrivers() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllDrivers")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_GetAllDrivers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllDrivers'
type HandlerFactory_GetAllDrivers_Call struct {
	*mock.Call
}

// GetAllDrivers is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) GetAllDrivers() *HandlerFactory_GetAllDrivers_Call {
	return &HandlerFactory_GetAllDrivers_Call{Call: _e.mock.On("GetAllDrivers")}
}

func (_c *HandlerFactory_GetAllDrivers_Call) Run(run func()) *HandlerFactory_GetAllDrivers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_GetAllDrivers_Call) Return(_a0 http.Handler) *HandlerFactory_GetAllDrivers_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_GetAllDrivers_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_GetAllDrivers_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllNotifications provides a mock function with given fields:
func (_m *HandlerFactory) GetAllNotifications() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllNotifications")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_GetAllNotifications_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllNotifications'
type HandlerFactory_GetAllNotifications_Call struct {
	*mock.Call
}

// GetAllNotifications is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) GetAllNotifications() *HandlerFactory_GetAllNotifications_Call {
	return &HandlerFactory_GetAllNotifications_Call{Call: _e.mock.On("GetAllNotifications")}
}

func (_c *HandlerFactory_GetAllNotifications_Call) Run(run func()) *HandlerFactory_GetAllNotifications_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_GetAllNotifications_Call) Return(_a0 http.Handler) *HandlerFactory_GetAllNotifications_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_GetAllNotifications_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_GetAllNotifications_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllParkingLots provides a mock function with given fields:
func (_m *HandlerFactory) GetAllParkingLots() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllParkingLots")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_GetAllParkingLots_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllParkingLots'
type HandlerFactory_GetAllParkingLots_Call struct {
	*mock.Call
}

// GetAllParkingLots is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) GetAllParkingLots() *HandlerFactory_GetAllParkingLots_Call {
	return &HandlerFactory_GetAllParkingLots_Call{Call: _e.mock.On("GetAllParkingLots")}
}

func (_c *HandlerFactory_GetAllParkingLots_Call) Run(run func()) *HandlerFactory_GetAllParkingLots_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_GetAllParkingLots_Call) Return(_a0 http.Handler) *HandlerFactory_GetAllParkingLots_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_GetAllParkingLots_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_GetAllParkingLots_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllParkingRequests provides a mock function with given fields:
func (_m *HandlerFactory) GetAllParkingRequests() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllParkingRequests")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_GetAllParkingRequests_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllParkingRequests'
type HandlerFactory_GetAllParkingRequests_Call struct {
	*mock.Call
}

// GetAllParkingRequests is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) GetAllParkingRequests() *HandlerFactory_GetAllParkingRequests_Call {
	return &HandlerFactory_GetAllParkingRequests_Call{Call: _e.mock.On("GetAllParkingRequests")}
}

func (_c *HandlerFactory_GetAllParkingRequests_Call) Run(run func()) *HandlerFactory_GetAllParkingRequests_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_GetAllParkingRequests_Call) Return(_a0 http.Handler) *HandlerFactory_GetAllParkingRequests_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_GetAllParkingRequests_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_GetAllParkingRequests_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllParkingRequestsForDriver provides a mock function with given fields:
func (_m *HandlerFactory) GetAllParkingRequestsForDriver() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllParkingRequestsForDriver")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_GetAllParkingRequestsForDriver_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllParkingRequestsForDriver'
type HandlerFactory_GetAllParkingRequestsForDriver_Call struct {
	*mock.Call
}

// GetAllParkingRequestsForDriver is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) GetAllParkingRequestsForDriver() *HandlerFactory_GetAllParkingRequestsForDriver_Call {
	return &HandlerFactory_GetAllParkingRequestsForDriver_Call{Call: _e.mock.On("GetAllParkingRequestsForDriver")}
}

func (_c *HandlerFactory_GetAllParkingRequestsForDriver_Call) Run(run func()) *HandlerFactory_GetAllParkingRequestsForDriver_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_GetAllParkingRequestsForDriver_Call) Return(_a0 http.Handler) *HandlerFactory_GetAllParkingRequestsForDriver_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_GetAllParkingRequestsForDriver_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_GetAllParkingRequestsForDriver_Call {
	_c.Call.Return(run)
	return _c
}

// GetSingleAlert provides a mock function with given fields:
func (_m *HandlerFactory) GetSingleAlert() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSingleAlert")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_GetSingleAlert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSingleAlert'
type HandlerFactory_GetSingleAlert_Call struct {
	*mock.Call
}

// GetSingleAlert is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) GetSingleAlert() *HandlerFactory_GetSingleAlert_Call {
	return &HandlerFactory_GetSingleAlert_Call{Call: _e.mock.On("GetSingleAlert")}
}

func (_c *HandlerFactory_GetSingleAlert_Call) Run(run func()) *HandlerFactory_GetSingleAlert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_GetSingleAlert_Call) Return(_a0 http.Handler) *HandlerFactory_GetSingleAlert_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_GetSingleAlert_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_GetSingleAlert_Call {
	_c.Call.Return(run)
	return _c
}

// GetSingleParkingLot provides a mock function with given fields:
func (_m *HandlerFactory) GetSingleParkingLot() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSingleParkingLot")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_GetSingleParkingLot_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSingleParkingLot'
type HandlerFactory_GetSingleParkingLot_Call struct {
	*mock.Call
}

// GetSingleParkingLot is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) GetSingleParkingLot() *HandlerFactory_GetSingleParkingLot_Call {
	return &HandlerFactory_GetSingleParkingLot_Call{Call: _e.mock.On("GetSingleParkingLot")}
}

func (_c *HandlerFactory_GetSingleParkingLot_Call) Run(run func()) *HandlerFactory_GetSingleParkingLot_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_GetSingleParkingLot_Call) Return(_a0 http.Handler) *HandlerFactory_GetSingleParkingLot_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_GetSingleParkingLot_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_GetSingleParkingLot_Call {
	_c.Call.Return(run)
	return _c
}

// GetSingleParkingSpace provides a mock function with given fields:
func (_m *HandlerFactory) GetSingleParkingSpace() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetSingleParkingSpace")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_GetSingleParkingSpace_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSingleParkingSpace'
type HandlerFactory_GetSingleParkingSpace_Call struct {
	*mock.Call
}

// GetSingleParkingSpace is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) GetSingleParkingSpace() *HandlerFactory_GetSingleParkingSpace_Call {
	return &HandlerFactory_GetSingleParkingSpace_Call{Call: _e.mock.On("GetSingleParkingSpace")}
}

func (_c *HandlerFactory_GetSingleParkingSpace_Call) Run(run func()) *HandlerFactory_GetSingleParkingSpace_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_GetSingleParkingSpace_Call) Return(_a0 http.Handler) *HandlerFactory_GetSingleParkingSpace_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_GetSingleParkingSpace_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_GetSingleParkingSpace_Call {
	_c.Call.Return(run)
	return _c
}

// ParkingLotCreate provides a mock function with given fields:
func (_m *HandlerFactory) ParkingLotCreate() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ParkingLotCreate")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_ParkingLotCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParkingLotCreate'
type HandlerFactory_ParkingLotCreate_Call struct {
	*mock.Call
}

// ParkingLotCreate is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) ParkingLotCreate() *HandlerFactory_ParkingLotCreate_Call {
	return &HandlerFactory_ParkingLotCreate_Call{Call: _e.mock.On("ParkingLotCreate")}
}

func (_c *HandlerFactory_ParkingLotCreate_Call) Run(run func()) *HandlerFactory_ParkingLotCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_ParkingLotCreate_Call) Return(_a0 http.Handler) *HandlerFactory_ParkingLotCreate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_ParkingLotCreate_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_ParkingLotCreate_Call {
	_c.Call.Return(run)
	return _c
}

// ParkingRequestCreate provides a mock function with given fields:
func (_m *HandlerFactory) ParkingRequestCreate() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ParkingRequestCreate")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_ParkingRequestCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParkingRequestCreate'
type HandlerFactory_ParkingRequestCreate_Call struct {
	*mock.Call
}

// ParkingRequestCreate is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) ParkingRequestCreate() *HandlerFactory_ParkingRequestCreate_Call {
	return &HandlerFactory_ParkingRequestCreate_Call{Call: _e.mock.On("ParkingRequestCreate")}
}

func (_c *HandlerFactory_ParkingRequestCreate_Call) Run(run func()) *HandlerFactory_ParkingRequestCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_ParkingRequestCreate_Call) Return(_a0 http.Handler) *HandlerFactory_ParkingRequestCreate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_ParkingRequestCreate_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_ParkingRequestCreate_Call {
	_c.Call.Return(run)
	return _c
}

// ParkingRequestStatusUpdate provides a mock function with given fields:
func (_m *HandlerFactory) ParkingRequestStatusUpdate() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ParkingRequestStatusUpdate")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_ParkingRequestStatusUpdate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ParkingRequestStatusUpdate'
type HandlerFactory_ParkingRequestStatusUpdate_Call struct {
	*mock.Call
}

// ParkingRequestStatusUpdate is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) ParkingRequestStatusUpdate() *HandlerFactory_ParkingRequestStatusUpdate_Call {
	return &HandlerFactory_ParkingRequestStatusUpdate_Call{Call: _e.mock.On("ParkingRequestStatusUpdate")}
}

func (_c *HandlerFactory_ParkingRequestStatusUpdate_Call) Run(run func()) *HandlerFactory_ParkingRequestStatusUpdate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_ParkingRequestStatusUpdate_Call) Return(_a0 http.Handler) *HandlerFactory_ParkingRequestStatusUpdate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_ParkingRequestStatusUpdate_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_ParkingRequestStatusUpdate_Call {
	_c.Call.Return(run)
	return _c
}

// PaymentCreate provides a mock function with given fields:
func (_m *HandlerFactory) PaymentCreate() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for PaymentCreate")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_PaymentCreate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PaymentCreate'
type HandlerFactory_PaymentCreate_Call struct {
	*mock.Call
}

// PaymentCreate is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) PaymentCreate() *HandlerFactory_PaymentCreate_Call {
	return &HandlerFactory_PaymentCreate_Call{Call: _e.mock.On("PaymentCreate")}
}

func (_c *HandlerFactory_PaymentCreate_Call) Run(run func()) *HandlerFactory_PaymentCreate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_PaymentCreate_Call) Return(_a0 http.Handler) *HandlerFactory_PaymentCreate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_PaymentCreate_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_PaymentCreate_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateParkingSpaceStatus provides a mock function with given fields:
func (_m *HandlerFactory) UpdateParkingSpaceStatus() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UpdateParkingSpaceStatus")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_UpdateParkingSpaceStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateParkingSpaceStatus'
type HandlerFactory_UpdateParkingSpaceStatus_Call struct {
	*mock.Call
}

// UpdateParkingSpaceStatus is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) UpdateParkingSpaceStatus() *HandlerFactory_UpdateParkingSpaceStatus_Call {
	return &HandlerFactory_UpdateParkingSpaceStatus_Call{Call: _e.mock.On("UpdateParkingSpaceStatus")}
}

func (_c *HandlerFactory_UpdateParkingSpaceStatus_Call) Run(run func()) *HandlerFactory_UpdateParkingSpaceStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_UpdateParkingSpaceStatus_Call) Return(_a0 http.Handler) *HandlerFactory_UpdateParkingSpaceStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_UpdateParkingSpaceStatus_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_UpdateParkingSpaceStatus_Call {
	_c.Call.Return(run)
	return _c
}

// UserAuthorise provides a mock function with given fields:
func (_m *HandlerFactory) UserAuthorise() http.Handler {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UserAuthorise")
	}

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// HandlerFactory_UserAuthorise_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserAuthorise'
type HandlerFactory_UserAuthorise_Call struct {
	*mock.Call
}

// UserAuthorise is a helper method to define mock.On call
func (_e *HandlerFactory_Expecter) UserAuthorise() *HandlerFactory_UserAuthorise_Call {
	return &HandlerFactory_UserAuthorise_Call{Call: _e.mock.On("UserAuthorise")}
}

func (_c *HandlerFactory_UserAuthorise_Call) Run(run func()) *HandlerFactory_UserAuthorise_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HandlerFactory_UserAuthorise_Call) Return(_a0 http.Handler) *HandlerFactory_UserAuthorise_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HandlerFactory_UserAuthorise_Call) RunAndReturn(run func() http.Handler) *HandlerFactory_UserAuthorise_Call {
	_c.Call.Return(run)
	return _c
}

// NewHandlerFactory creates a new instance of HandlerFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandlerFactory(t interface {
	mock.TestingT
	Cleanup(func())
}) *HandlerFactory {
	mock := &HandlerFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
