// Code generated by mockery v2.46.3. DO NOT EDIT.

package service

import (
	model "github.com/iqunlim/easyblog/model"
	mock "github.com/stretchr/testify/mock"
)

// MockUserService is an autogenerated mock type for the UserService type
type MockUserService struct {
	mock.Mock
}

type MockUserService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserService) EXPECT() *MockUserService_Expecter {
	return &MockUserService_Expecter{mock: &_m.Mock}
}

// Register provides a mock function with given fields: u
func (_m *MockUserService) Register(u *model.User) error {
	ret := _m.Called(u)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserService_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type MockUserService_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - u *model.User
func (_e *MockUserService_Expecter) Register(u interface{}) *MockUserService_Register_Call {
	return &MockUserService_Register_Call{Call: _e.mock.On("Register", u)}
}

func (_c *MockUserService_Register_Call) Run(run func(u *model.User)) *MockUserService_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.User))
	})
	return _c
}

func (_c *MockUserService_Register_Call) Return(_a0 error) *MockUserService_Register_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserService_Register_Call) RunAndReturn(run func(*model.User) error) *MockUserService_Register_Call {
	_c.Call.Return(run)
	return _c
}

// Verify provides a mock function with given fields: u
func (_m *MockUserService) Verify(u *model.User) (*model.User, error) {
	ret := _m.Called(u)

	if len(ret) == 0 {
		panic("no return value specified for Verify")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.User) (*model.User, error)); ok {
		return rf(u)
	}
	if rf, ok := ret.Get(0).(func(*model.User) *model.User); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.User) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserService_Verify_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Verify'
type MockUserService_Verify_Call struct {
	*mock.Call
}

// Verify is a helper method to define mock.On call
//   - u *model.User
func (_e *MockUserService_Expecter) Verify(u interface{}) *MockUserService_Verify_Call {
	return &MockUserService_Verify_Call{Call: _e.mock.On("Verify", u)}
}

func (_c *MockUserService_Verify_Call) Run(run func(u *model.User)) *MockUserService_Verify_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.User))
	})
	return _c
}

func (_c *MockUserService_Verify_Call) Return(_a0 *model.User, _a1 error) *MockUserService_Verify_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserService_Verify_Call) RunAndReturn(run func(*model.User) (*model.User, error)) *MockUserService_Verify_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserService creates a new instance of MockUserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserService {
	mock := &MockUserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
