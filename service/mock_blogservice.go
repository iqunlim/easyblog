// Code generated by mockery v2.50.2. DO NOT EDIT.

package service

import (
	context "context"

	model "github.com/iqunlim/easyblog/model"
	mock "github.com/stretchr/testify/mock"
)

// MockBlogService is an autogenerated mock type for the BlogService type
type MockBlogService struct {
	mock.Mock
}

type MockBlogService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBlogService) EXPECT() *MockBlogService_Expecter {
	return &MockBlogService_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, id
func (_m *MockBlogService) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBlogService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockBlogService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockBlogService_Expecter) Delete(ctx interface{}, id interface{}) *MockBlogService_Delete_Call {
	return &MockBlogService_Delete_Call{Call: _e.mock.On("Delete", ctx, id)}
}

func (_c *MockBlogService_Delete_Call) Run(run func(ctx context.Context, id string)) *MockBlogService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockBlogService_Delete_Call) Return(_a0 error) *MockBlogService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBlogService_Delete_Call) RunAndReturn(run func(context.Context, string) error) *MockBlogService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// GetAll provides a mock function with given fields: ctx, params, htmlformat
func (_m *MockBlogService) GetAll(ctx context.Context, params string, htmlformat bool) ([]*model.BlogPost, error) {
	ret := _m.Called(ctx, params, htmlformat)

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []*model.BlogPost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) ([]*model.BlogPost, error)); ok {
		return rf(ctx, params, htmlformat)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) []*model.BlogPost); ok {
		r0 = rf(ctx, params, htmlformat)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.BlogPost)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, params, htmlformat)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBlogService_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type MockBlogService_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
//   - ctx context.Context
//   - params string
//   - htmlformat bool
func (_e *MockBlogService_Expecter) GetAll(ctx interface{}, params interface{}, htmlformat interface{}) *MockBlogService_GetAll_Call {
	return &MockBlogService_GetAll_Call{Call: _e.mock.On("GetAll", ctx, params, htmlformat)}
}

func (_c *MockBlogService_GetAll_Call) Run(run func(ctx context.Context, params string, htmlformat bool)) *MockBlogService_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(bool))
	})
	return _c
}

func (_c *MockBlogService_GetAll_Call) Return(_a0 []*model.BlogPost, _a1 error) *MockBlogService_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBlogService_GetAll_Call) RunAndReturn(run func(context.Context, string, bool) ([]*model.BlogPost, error)) *MockBlogService_GetAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllNoContent provides a mock function with given fields: ctx
func (_m *MockBlogService) GetAllNoContent(ctx context.Context) ([]*model.BlogPost, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllNoContent")
	}

	var r0 []*model.BlogPost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*model.BlogPost, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*model.BlogPost); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.BlogPost)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBlogService_GetAllNoContent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllNoContent'
type MockBlogService_GetAllNoContent_Call struct {
	*mock.Call
}

// GetAllNoContent is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockBlogService_Expecter) GetAllNoContent(ctx interface{}) *MockBlogService_GetAllNoContent_Call {
	return &MockBlogService_GetAllNoContent_Call{Call: _e.mock.On("GetAllNoContent", ctx)}
}

func (_c *MockBlogService_GetAllNoContent_Call) Run(run func(ctx context.Context)) *MockBlogService_GetAllNoContent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockBlogService_GetAllNoContent_Call) Return(_a0 []*model.BlogPost, _a1 error) *MockBlogService_GetAllNoContent_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBlogService_GetAllNoContent_Call) RunAndReturn(run func(context.Context) ([]*model.BlogPost, error)) *MockBlogService_GetAllNoContent_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, id, htmlformat
func (_m *MockBlogService) GetByID(ctx context.Context, id string, htmlformat bool) (*model.BlogPost, error) {
	ret := _m.Called(ctx, id, htmlformat)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *model.BlogPost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) (*model.BlogPost, error)); ok {
		return rf(ctx, id, htmlformat)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) *model.BlogPost); ok {
		r0 = rf(ctx, id, htmlformat)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.BlogPost)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, id, htmlformat)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBlogService_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type MockBlogService_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - htmlformat bool
func (_e *MockBlogService_Expecter) GetByID(ctx interface{}, id interface{}, htmlformat interface{}) *MockBlogService_GetByID_Call {
	return &MockBlogService_GetByID_Call{Call: _e.mock.On("GetByID", ctx, id, htmlformat)}
}

func (_c *MockBlogService_GetByID_Call) Run(run func(ctx context.Context, id string, htmlformat bool)) *MockBlogService_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(bool))
	})
	return _c
}

func (_c *MockBlogService_GetByID_Call) Return(_a0 *model.BlogPost, _a1 error) *MockBlogService_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBlogService_GetByID_Call) RunAndReturn(run func(context.Context, string, bool) (*model.BlogPost, error)) *MockBlogService_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// Post provides a mock function with given fields: ctx, blog
func (_m *MockBlogService) Post(ctx context.Context, blog *model.BlogPost) (*model.BlogPost, error) {
	ret := _m.Called(ctx, blog)

	if len(ret) == 0 {
		panic("no return value specified for Post")
	}

	var r0 *model.BlogPost
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.BlogPost) (*model.BlogPost, error)); ok {
		return rf(ctx, blog)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.BlogPost) *model.BlogPost); ok {
		r0 = rf(ctx, blog)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.BlogPost)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.BlogPost) error); ok {
		r1 = rf(ctx, blog)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockBlogService_Post_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Post'
type MockBlogService_Post_Call struct {
	*mock.Call
}

// Post is a helper method to define mock.On call
//   - ctx context.Context
//   - blog *model.BlogPost
func (_e *MockBlogService_Expecter) Post(ctx interface{}, blog interface{}) *MockBlogService_Post_Call {
	return &MockBlogService_Post_Call{Call: _e.mock.On("Post", ctx, blog)}
}

func (_c *MockBlogService_Post_Call) Run(run func(ctx context.Context, blog *model.BlogPost)) *MockBlogService_Post_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*model.BlogPost))
	})
	return _c
}

func (_c *MockBlogService_Post_Call) Return(_a0 *model.BlogPost, _a1 error) *MockBlogService_Post_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockBlogService_Post_Call) RunAndReturn(run func(context.Context, *model.BlogPost) (*model.BlogPost, error)) *MockBlogService_Post_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, id, blog
func (_m *MockBlogService) Update(ctx context.Context, id string, blog *model.BlogPost) error {
	ret := _m.Called(ctx, id, blog)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.BlogPost) error); ok {
		r0 = rf(ctx, id, blog)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockBlogService_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockBlogService_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
//   - blog *model.BlogPost
func (_e *MockBlogService_Expecter) Update(ctx interface{}, id interface{}, blog interface{}) *MockBlogService_Update_Call {
	return &MockBlogService_Update_Call{Call: _e.mock.On("Update", ctx, id, blog)}
}

func (_c *MockBlogService_Update_Call) Run(run func(ctx context.Context, id string, blog *model.BlogPost)) *MockBlogService_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(*model.BlogPost))
	})
	return _c
}

func (_c *MockBlogService_Update_Call) Return(_a0 error) *MockBlogService_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockBlogService_Update_Call) RunAndReturn(run func(context.Context, string, *model.BlogPost) error) *MockBlogService_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBlogService creates a new instance of MockBlogService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBlogService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBlogService {
	mock := &MockBlogService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
