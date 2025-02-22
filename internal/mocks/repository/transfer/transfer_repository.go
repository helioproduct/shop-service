// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"
	domain "shop-service/internal/domain"

	mock "github.com/stretchr/testify/mock"

	repositorytransfer "shop-service/internal/repository/transfer"
)

// TransferRepository is an autogenerated mock type for the TransferRepository type
type TransferRepository struct {
	mock.Mock
}

type TransferRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *TransferRepository) EXPECT() *TransferRepository_Expecter {
	return &TransferRepository_Expecter{mock: &_m.Mock}
}

// CreateTransfer provides a mock function with given fields: ctx, _a1
func (_m *TransferRepository) CreateTransfer(ctx context.Context, _a1 domain.Transfer) (*domain.Transfer, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateTransfer")
	}

	var r0 *domain.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Transfer) (*domain.Transfer, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.Transfer) *domain.Transfer); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Transfer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.Transfer) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransferRepository_CreateTransfer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTransfer'
type TransferRepository_CreateTransfer_Call struct {
	*mock.Call
}

// CreateTransfer is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 domain.Transfer
func (_e *TransferRepository_Expecter) CreateTransfer(ctx interface{}, _a1 interface{}) *TransferRepository_CreateTransfer_Call {
	return &TransferRepository_CreateTransfer_Call{Call: _e.mock.On("CreateTransfer", ctx, _a1)}
}

func (_c *TransferRepository_CreateTransfer_Call) Run(run func(ctx context.Context, _a1 domain.Transfer)) *TransferRepository_CreateTransfer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(domain.Transfer))
	})
	return _c
}

func (_c *TransferRepository_CreateTransfer_Call) Return(_a0 *domain.Transfer, _a1 error) *TransferRepository_CreateTransfer_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TransferRepository_CreateTransfer_Call) RunAndReturn(run func(context.Context, domain.Transfer) (*domain.Transfer, error)) *TransferRepository_CreateTransfer_Call {
	_c.Call.Return(run)
	return _c
}

// GetReceivedCoinsSummary provides a mock function with given fields: ctx, toUsername
func (_m *TransferRepository) GetReceivedCoinsSummary(ctx context.Context, toUsername string) ([]*repositorytransfer.ReceivedCoinsSummary, error) {
	ret := _m.Called(ctx, toUsername)

	if len(ret) == 0 {
		panic("no return value specified for GetReceivedCoinsSummary")
	}

	var r0 []*repositorytransfer.ReceivedCoinsSummary
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*repositorytransfer.ReceivedCoinsSummary, error)); ok {
		return rf(ctx, toUsername)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*repositorytransfer.ReceivedCoinsSummary); ok {
		r0 = rf(ctx, toUsername)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repositorytransfer.ReceivedCoinsSummary)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, toUsername)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransferRepository_GetReceivedCoinsSummary_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReceivedCoinsSummary'
type TransferRepository_GetReceivedCoinsSummary_Call struct {
	*mock.Call
}

// GetReceivedCoinsSummary is a helper method to define mock.On call
//   - ctx context.Context
//   - toUsername string
func (_e *TransferRepository_Expecter) GetReceivedCoinsSummary(ctx interface{}, toUsername interface{}) *TransferRepository_GetReceivedCoinsSummary_Call {
	return &TransferRepository_GetReceivedCoinsSummary_Call{Call: _e.mock.On("GetReceivedCoinsSummary", ctx, toUsername)}
}

func (_c *TransferRepository_GetReceivedCoinsSummary_Call) Run(run func(ctx context.Context, toUsername string)) *TransferRepository_GetReceivedCoinsSummary_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *TransferRepository_GetReceivedCoinsSummary_Call) Return(_a0 []*repositorytransfer.ReceivedCoinsSummary, _a1 error) *TransferRepository_GetReceivedCoinsSummary_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TransferRepository_GetReceivedCoinsSummary_Call) RunAndReturn(run func(context.Context, string) ([]*repositorytransfer.ReceivedCoinsSummary, error)) *TransferRepository_GetReceivedCoinsSummary_Call {
	_c.Call.Return(run)
	return _c
}

// GetSentCoinsSummary provides a mock function with given fields: ctx, fromUsername
func (_m *TransferRepository) GetSentCoinsSummary(ctx context.Context, fromUsername string) ([]*repositorytransfer.SentCoinsSummary, error) {
	ret := _m.Called(ctx, fromUsername)

	if len(ret) == 0 {
		panic("no return value specified for GetSentCoinsSummary")
	}

	var r0 []*repositorytransfer.SentCoinsSummary
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*repositorytransfer.SentCoinsSummary, error)); ok {
		return rf(ctx, fromUsername)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*repositorytransfer.SentCoinsSummary); ok {
		r0 = rf(ctx, fromUsername)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*repositorytransfer.SentCoinsSummary)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, fromUsername)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransferRepository_GetSentCoinsSummary_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSentCoinsSummary'
type TransferRepository_GetSentCoinsSummary_Call struct {
	*mock.Call
}

// GetSentCoinsSummary is a helper method to define mock.On call
//   - ctx context.Context
//   - fromUsername string
func (_e *TransferRepository_Expecter) GetSentCoinsSummary(ctx interface{}, fromUsername interface{}) *TransferRepository_GetSentCoinsSummary_Call {
	return &TransferRepository_GetSentCoinsSummary_Call{Call: _e.mock.On("GetSentCoinsSummary", ctx, fromUsername)}
}

func (_c *TransferRepository_GetSentCoinsSummary_Call) Run(run func(ctx context.Context, fromUsername string)) *TransferRepository_GetSentCoinsSummary_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *TransferRepository_GetSentCoinsSummary_Call) Return(_a0 []*repositorytransfer.SentCoinsSummary, _a1 error) *TransferRepository_GetSentCoinsSummary_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TransferRepository_GetSentCoinsSummary_Call) RunAndReturn(run func(context.Context, string) ([]*repositorytransfer.SentCoinsSummary, error)) *TransferRepository_GetSentCoinsSummary_Call {
	_c.Call.Return(run)
	return _c
}

// NewTransferRepository creates a new instance of TransferRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTransferRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TransferRepository {
	mock := &TransferRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
