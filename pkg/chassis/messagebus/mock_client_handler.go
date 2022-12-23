// Code generated by mockery v2.15.0. DO NOT EDIT.

package messagebus

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockClientHandler is an autogenerated mock type for the ClientHandler type
type MockClientHandler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, msg
func (_m *MockClientHandler) Handle(ctx context.Context, msg Message) error {
	ret := _m.Called(ctx, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, Message) error); ok {
		r0 = rf(ctx, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockClientHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockClientHandler creates a new instance of MockClientHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockClientHandler(t mockConstructorTestingTNewMockClientHandler) *MockClientHandler {
	mock := &MockClientHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
