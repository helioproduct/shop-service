// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// GetBalance provides a mock function with given fields: ctx, username
func (_m *UserUsecase) GetBalance(ctx context.Context, username string) (uint64, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for GetBalance")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (uint64, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) uint64); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
