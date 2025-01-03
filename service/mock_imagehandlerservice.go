// Code generated by mockery v2.50.2. DO NOT EDIT.

package service

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"
)

// MockImageHandlerService is an autogenerated mock type for the ImageHandlerService type
type MockImageHandlerService struct {
	mock.Mock
}

type MockImageHandlerService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockImageHandlerService) EXPECT() *MockImageHandlerService_Expecter {
	return &MockImageHandlerService_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: ctx, fileName
func (_m *MockImageHandlerService) Delete(ctx context.Context, fileName string) error {
	ret := _m.Called(ctx, fileName)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, fileName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockImageHandlerService_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockImageHandlerService_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - fileName string
func (_e *MockImageHandlerService_Expecter) Delete(ctx interface{}, fileName interface{}) *MockImageHandlerService_Delete_Call {
	return &MockImageHandlerService_Delete_Call{Call: _e.mock.On("Delete", ctx, fileName)}
}

func (_c *MockImageHandlerService_Delete_Call) Run(run func(ctx context.Context, fileName string)) *MockImageHandlerService_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockImageHandlerService_Delete_Call) Return(_a0 error) *MockImageHandlerService_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockImageHandlerService_Delete_Call) RunAndReturn(run func(context.Context, string) error) *MockImageHandlerService_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Download provides a mock function with given fields: ctx, fileName
func (_m *MockImageHandlerService) Download(ctx context.Context, fileName string) (io.ReadCloser, error) {
	ret := _m.Called(ctx, fileName)

	if len(ret) == 0 {
		panic("no return value specified for Download")
	}

	var r0 io.ReadCloser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (io.ReadCloser, error)); ok {
		return rf(ctx, fileName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) io.ReadCloser); ok {
		r0 = rf(ctx, fileName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, fileName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockImageHandlerService_Download_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Download'
type MockImageHandlerService_Download_Call struct {
	*mock.Call
}

// Download is a helper method to define mock.On call
//   - ctx context.Context
//   - fileName string
func (_e *MockImageHandlerService_Expecter) Download(ctx interface{}, fileName interface{}) *MockImageHandlerService_Download_Call {
	return &MockImageHandlerService_Download_Call{Call: _e.mock.On("Download", ctx, fileName)}
}

func (_c *MockImageHandlerService_Download_Call) Run(run func(ctx context.Context, fileName string)) *MockImageHandlerService_Download_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockImageHandlerService_Download_Call) Return(_a0 io.ReadCloser, _a1 error) *MockImageHandlerService_Download_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockImageHandlerService_Download_Call) RunAndReturn(run func(context.Context, string) (io.ReadCloser, error)) *MockImageHandlerService_Download_Call {
	_c.Call.Return(run)
	return _c
}

// Upload provides a mock function with given fields: ctx, fileBody, fileReader
func (_m *MockImageHandlerService) Upload(ctx context.Context, fileBody io.Reader, fileReader *multipart.FileHeader) (string, error) {
	ret := _m.Called(ctx, fileBody, fileReader)

	if len(ret) == 0 {
		panic("no return value specified for Upload")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader, *multipart.FileHeader) (string, error)); ok {
		return rf(ctx, fileBody, fileReader)
	}
	if rf, ok := ret.Get(0).(func(context.Context, io.Reader, *multipart.FileHeader) string); ok {
		r0 = rf(ctx, fileBody, fileReader)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, io.Reader, *multipart.FileHeader) error); ok {
		r1 = rf(ctx, fileBody, fileReader)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockImageHandlerService_Upload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upload'
type MockImageHandlerService_Upload_Call struct {
	*mock.Call
}

// Upload is a helper method to define mock.On call
//   - ctx context.Context
//   - fileBody io.Reader
//   - fileReader *multipart.FileHeader
func (_e *MockImageHandlerService_Expecter) Upload(ctx interface{}, fileBody interface{}, fileReader interface{}) *MockImageHandlerService_Upload_Call {
	return &MockImageHandlerService_Upload_Call{Call: _e.mock.On("Upload", ctx, fileBody, fileReader)}
}

func (_c *MockImageHandlerService_Upload_Call) Run(run func(ctx context.Context, fileBody io.Reader, fileReader *multipart.FileHeader)) *MockImageHandlerService_Upload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(io.Reader), args[2].(*multipart.FileHeader))
	})
	return _c
}

func (_c *MockImageHandlerService_Upload_Call) Return(_a0 string, _a1 error) *MockImageHandlerService_Upload_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockImageHandlerService_Upload_Call) RunAndReturn(run func(context.Context, io.Reader, *multipart.FileHeader) (string, error)) *MockImageHandlerService_Upload_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockImageHandlerService creates a new instance of MockImageHandlerService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockImageHandlerService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockImageHandlerService {
	mock := &MockImageHandlerService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}