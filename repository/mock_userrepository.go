// Code generated by mockery v2.46.3. DO NOT EDIT.

package repository

import (
	model "github.com/iqunlim/easyblog/model"
	mock "github.com/stretchr/testify/mock"
)

// MockUserRepository is an autogenerated mock type for the UserRepository type
type MockUserRepository struct {
	mock.Mock
}

type MockUserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserRepository) EXPECT() *MockUserRepository_Expecter {
	return &MockUserRepository_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: u
func (_m *MockUserRepository) Create(u *model.User) error {
	ret := _m.Called(u)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
		r0 = rf(u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockUserRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - u *model.User
func (_e *MockUserRepository_Expecter) Create(u interface{}) *MockUserRepository_Create_Call {
	return &MockUserRepository_Create_Call{Call: _e.mock.On("Create", u)}
}

func (_c *MockUserRepository_Create_Call) Run(run func(u *model.User)) *MockUserRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.User))
	})
	return _c
}

func (_c *MockUserRepository_Create_Call) Return(_a0 error) *MockUserRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserRepository_Create_Call) RunAndReturn(run func(*model.User) error) *MockUserRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// GetByUsername provides a mock function with given fields: username
func (_m *MockUserRepository) GetByUsername(username string) (*model.User, error) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for GetByUsername")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserRepository_GetByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByUsername'
type MockUserRepository_GetByUsername_Call struct {
	*mock.Call
}

// GetByUsername is a helper method to define mock.On call
//   - username string
func (_e *MockUserRepository_Expecter) GetByUsername(username interface{}) *MockUserRepository_GetByUsername_Call {
	return &MockUserRepository_GetByUsername_Call{Call: _e.mock.On("GetByUsername", username)}
}

func (_c *MockUserRepository_GetByUsername_Call) Run(run func(username string)) *MockUserRepository_GetByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockUserRepository_GetByUsername_Call) Return(_a0 *model.User, _a1 error) *MockUserRepository_GetByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserRepository_GetByUsername_Call) RunAndReturn(run func(string) (*model.User, error)) *MockUserRepository_GetByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserRepository creates a new instance of MockUserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserRepository {
	mock := &MockUserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
