// Code generated by mockery v2.32.4. DO NOT EDIT.

package mockbind

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	context "context"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// ContractBackend is an autogenerated mock type for the ContractBackend type
type ContractBackend struct {
	mock.Mock
}

type ContractBackend_Expecter struct {
	mock *mock.Mock
}

func (_m *ContractBackend) EXPECT() *ContractBackend_Expecter {
	return &ContractBackend_Expecter{mock: &_m.Mock}
}

// CallContract provides a mock function with given fields: ctx, call, blockNumber
func (_m *ContractBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, call, blockNumber)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg, *big.Int) ([]byte, error)); ok {
		return rf(ctx, call, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg, *big.Int) []byte); ok {
		r0 = rf(ctx, call, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.CallMsg, *big.Int) error); ok {
		r1 = rf(ctx, call, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContractBackend_CallContract_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CallContract'
type ContractBackend_CallContract_Call struct {
	*mock.Call
}

// CallContract is a helper method to define mock.On call
//   - ctx context.Context
//   - call ethereum.CallMsg
//   - blockNumber *big.Int
func (_e *ContractBackend_Expecter) CallContract(ctx interface{}, call interface{}, blockNumber interface{}) *ContractBackend_CallContract_Call {
	return &ContractBackend_CallContract_Call{Call: _e.mock.On("CallContract", ctx, call, blockNumber)}
}

func (_c *ContractBackend_CallContract_Call) Run(run func(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int)) *ContractBackend_CallContract_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ethereum.CallMsg), args[2].(*big.Int))
	})
	return _c
}

func (_c *ContractBackend_CallContract_Call) Return(_a0 []byte, _a1 error) *ContractBackend_CallContract_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ContractBackend_CallContract_Call) RunAndReturn(run func(context.Context, ethereum.CallMsg, *big.Int) ([]byte, error)) *ContractBackend_CallContract_Call {
	_c.Call.Return(run)
	return _c
}

// CodeAt provides a mock function with given fields: ctx, contract, blockNumber
func (_m *ContractBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	ret := _m.Called(ctx, contract, blockNumber)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) ([]byte, error)); ok {
		return rf(ctx, contract, blockNumber)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address, *big.Int) []byte); ok {
		r0 = rf(ctx, contract, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(ctx, contract, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContractBackend_CodeAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CodeAt'
type ContractBackend_CodeAt_Call struct {
	*mock.Call
}

// CodeAt is a helper method to define mock.On call
//   - ctx context.Context
//   - contract common.Address
//   - blockNumber *big.Int
func (_e *ContractBackend_Expecter) CodeAt(ctx interface{}, contract interface{}, blockNumber interface{}) *ContractBackend_CodeAt_Call {
	return &ContractBackend_CodeAt_Call{Call: _e.mock.On("CodeAt", ctx, contract, blockNumber)}
}

func (_c *ContractBackend_CodeAt_Call) Run(run func(ctx context.Context, contract common.Address, blockNumber *big.Int)) *ContractBackend_CodeAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address), args[2].(*big.Int))
	})
	return _c
}

func (_c *ContractBackend_CodeAt_Call) Return(_a0 []byte, _a1 error) *ContractBackend_CodeAt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ContractBackend_CodeAt_Call) RunAndReturn(run func(context.Context, common.Address, *big.Int) ([]byte, error)) *ContractBackend_CodeAt_Call {
	_c.Call.Return(run)
	return _c
}

// EstimateGas provides a mock function with given fields: ctx, call
func (_m *ContractBackend) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	ret := _m.Called(ctx, call)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg) (uint64, error)); ok {
		return rf(ctx, call)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.CallMsg) uint64); ok {
		r0 = rf(ctx, call)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.CallMsg) error); ok {
		r1 = rf(ctx, call)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContractBackend_EstimateGas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EstimateGas'
type ContractBackend_EstimateGas_Call struct {
	*mock.Call
}

// EstimateGas is a helper method to define mock.On call
//   - ctx context.Context
//   - call ethereum.CallMsg
func (_e *ContractBackend_Expecter) EstimateGas(ctx interface{}, call interface{}) *ContractBackend_EstimateGas_Call {
	return &ContractBackend_EstimateGas_Call{Call: _e.mock.On("EstimateGas", ctx, call)}
}

func (_c *ContractBackend_EstimateGas_Call) Run(run func(ctx context.Context, call ethereum.CallMsg)) *ContractBackend_EstimateGas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ethereum.CallMsg))
	})
	return _c
}

func (_c *ContractBackend_EstimateGas_Call) Return(gas uint64, err error) *ContractBackend_EstimateGas_Call {
	_c.Call.Return(gas, err)
	return _c
}

func (_c *ContractBackend_EstimateGas_Call) RunAndReturn(run func(context.Context, ethereum.CallMsg) (uint64, error)) *ContractBackend_EstimateGas_Call {
	_c.Call.Return(run)
	return _c
}

// FilterLogs provides a mock function with given fields: ctx, query
func (_m *ContractBackend) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	ret := _m.Called(ctx, query)

	var r0 []types.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery) ([]types.Log, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery) []types.Log); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.FilterQuery) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContractBackend_FilterLogs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FilterLogs'
type ContractBackend_FilterLogs_Call struct {
	*mock.Call
}

// FilterLogs is a helper method to define mock.On call
//   - ctx context.Context
//   - query ethereum.FilterQuery
func (_e *ContractBackend_Expecter) FilterLogs(ctx interface{}, query interface{}) *ContractBackend_FilterLogs_Call {
	return &ContractBackend_FilterLogs_Call{Call: _e.mock.On("FilterLogs", ctx, query)}
}

func (_c *ContractBackend_FilterLogs_Call) Run(run func(ctx context.Context, query ethereum.FilterQuery)) *ContractBackend_FilterLogs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ethereum.FilterQuery))
	})
	return _c
}

func (_c *ContractBackend_FilterLogs_Call) Return(_a0 []types.Log, _a1 error) *ContractBackend_FilterLogs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ContractBackend_FilterLogs_Call) RunAndReturn(run func(context.Context, ethereum.FilterQuery) ([]types.Log, error)) *ContractBackend_FilterLogs_Call {
	_c.Call.Return(run)
	return _c
}

// HeaderByNumber provides a mock function with given fields: ctx, number
func (_m *ContractBackend) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	ret := _m.Called(ctx, number)

	var r0 *types.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) (*types.Header, error)); ok {
		return rf(ctx, number)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *big.Int) *types.Header); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *big.Int) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContractBackend_HeaderByNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HeaderByNumber'
type ContractBackend_HeaderByNumber_Call struct {
	*mock.Call
}

// HeaderByNumber is a helper method to define mock.On call
//   - ctx context.Context
//   - number *big.Int
func (_e *ContractBackend_Expecter) HeaderByNumber(ctx interface{}, number interface{}) *ContractBackend_HeaderByNumber_Call {
	return &ContractBackend_HeaderByNumber_Call{Call: _e.mock.On("HeaderByNumber", ctx, number)}
}

func (_c *ContractBackend_HeaderByNumber_Call) Run(run func(ctx context.Context, number *big.Int)) *ContractBackend_HeaderByNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*big.Int))
	})
	return _c
}

func (_c *ContractBackend_HeaderByNumber_Call) Return(_a0 *types.Header, _a1 error) *ContractBackend_HeaderByNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ContractBackend_HeaderByNumber_Call) RunAndReturn(run func(context.Context, *big.Int) (*types.Header, error)) *ContractBackend_HeaderByNumber_Call {
	_c.Call.Return(run)
	return _c
}

// PendingCodeAt provides a mock function with given fields: ctx, account
func (_m *ContractBackend) PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error) {
	ret := _m.Called(ctx, account)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) ([]byte, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) []byte); ok {
		r0 = rf(ctx, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContractBackend_PendingCodeAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PendingCodeAt'
type ContractBackend_PendingCodeAt_Call struct {
	*mock.Call
}

// PendingCodeAt is a helper method to define mock.On call
//   - ctx context.Context
//   - account common.Address
func (_e *ContractBackend_Expecter) PendingCodeAt(ctx interface{}, account interface{}) *ContractBackend_PendingCodeAt_Call {
	return &ContractBackend_PendingCodeAt_Call{Call: _e.mock.On("PendingCodeAt", ctx, account)}
}

func (_c *ContractBackend_PendingCodeAt_Call) Run(run func(ctx context.Context, account common.Address)) *ContractBackend_PendingCodeAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address))
	})
	return _c
}

func (_c *ContractBackend_PendingCodeAt_Call) Return(_a0 []byte, _a1 error) *ContractBackend_PendingCodeAt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ContractBackend_PendingCodeAt_Call) RunAndReturn(run func(context.Context, common.Address) ([]byte, error)) *ContractBackend_PendingCodeAt_Call {
	_c.Call.Return(run)
	return _c
}

// PendingNonceAt provides a mock function with given fields: ctx, account
func (_m *ContractBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	ret := _m.Called(ctx, account)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) (uint64, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, common.Address) uint64); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, common.Address) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContractBackend_PendingNonceAt_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PendingNonceAt'
type ContractBackend_PendingNonceAt_Call struct {
	*mock.Call
}

// PendingNonceAt is a helper method to define mock.On call
//   - ctx context.Context
//   - account common.Address
func (_e *ContractBackend_Expecter) PendingNonceAt(ctx interface{}, account interface{}) *ContractBackend_PendingNonceAt_Call {
	return &ContractBackend_PendingNonceAt_Call{Call: _e.mock.On("PendingNonceAt", ctx, account)}
}

func (_c *ContractBackend_PendingNonceAt_Call) Run(run func(ctx context.Context, account common.Address)) *ContractBackend_PendingNonceAt_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(common.Address))
	})
	return _c
}

func (_c *ContractBackend_PendingNonceAt_Call) Return(_a0 uint64, _a1 error) *ContractBackend_PendingNonceAt_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ContractBackend_PendingNonceAt_Call) RunAndReturn(run func(context.Context, common.Address) (uint64, error)) *ContractBackend_PendingNonceAt_Call {
	_c.Call.Return(run)
	return _c
}

// SendTransaction provides a mock function with given fields: ctx, tx
func (_m *ContractBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	ret := _m.Called(ctx, tx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Transaction) error); ok {
		r0 = rf(ctx, tx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ContractBackend_SendTransaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendTransaction'
type ContractBackend_SendTransaction_Call struct {
	*mock.Call
}

// SendTransaction is a helper method to define mock.On call
//   - ctx context.Context
//   - tx *types.Transaction
func (_e *ContractBackend_Expecter) SendTransaction(ctx interface{}, tx interface{}) *ContractBackend_SendTransaction_Call {
	return &ContractBackend_SendTransaction_Call{Call: _e.mock.On("SendTransaction", ctx, tx)}
}

func (_c *ContractBackend_SendTransaction_Call) Run(run func(ctx context.Context, tx *types.Transaction)) *ContractBackend_SendTransaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*types.Transaction))
	})
	return _c
}

func (_c *ContractBackend_SendTransaction_Call) Return(_a0 error) *ContractBackend_SendTransaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ContractBackend_SendTransaction_Call) RunAndReturn(run func(context.Context, *types.Transaction) error) *ContractBackend_SendTransaction_Call {
	_c.Call.Return(run)
	return _c
}

// SubscribeFilterLogs provides a mock function with given fields: ctx, query, ch
func (_m *ContractBackend) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	ret := _m.Called(ctx, query, ch)

	var r0 ethereum.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) (ethereum.Subscription, error)); ok {
		return rf(ctx, query, ch)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) ethereum.Subscription); ok {
		r0 = rf(ctx, query, ch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ethereum.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) error); ok {
		r1 = rf(ctx, query, ch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContractBackend_SubscribeFilterLogs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeFilterLogs'
type ContractBackend_SubscribeFilterLogs_Call struct {
	*mock.Call
}

// SubscribeFilterLogs is a helper method to define mock.On call
//   - ctx context.Context
//   - query ethereum.FilterQuery
//   - ch chan<- types.Log
func (_e *ContractBackend_Expecter) SubscribeFilterLogs(ctx interface{}, query interface{}, ch interface{}) *ContractBackend_SubscribeFilterLogs_Call {
	return &ContractBackend_SubscribeFilterLogs_Call{Call: _e.mock.On("SubscribeFilterLogs", ctx, query, ch)}
}

func (_c *ContractBackend_SubscribeFilterLogs_Call) Run(run func(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log)) *ContractBackend_SubscribeFilterLogs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ethereum.FilterQuery), args[2].(chan<- types.Log))
	})
	return _c
}

func (_c *ContractBackend_SubscribeFilterLogs_Call) Return(_a0 ethereum.Subscription, _a1 error) *ContractBackend_SubscribeFilterLogs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ContractBackend_SubscribeFilterLogs_Call) RunAndReturn(run func(context.Context, ethereum.FilterQuery, chan<- types.Log) (ethereum.Subscription, error)) *ContractBackend_SubscribeFilterLogs_Call {
	_c.Call.Return(run)
	return _c
}

// SuggestGasPrice provides a mock function with given fields: ctx
func (_m *ContractBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContractBackend_SuggestGasPrice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SuggestGasPrice'
type ContractBackend_SuggestGasPrice_Call struct {
	*mock.Call
}

// SuggestGasPrice is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ContractBackend_Expecter) SuggestGasPrice(ctx interface{}) *ContractBackend_SuggestGasPrice_Call {
	return &ContractBackend_SuggestGasPrice_Call{Call: _e.mock.On("SuggestGasPrice", ctx)}
}

func (_c *ContractBackend_SuggestGasPrice_Call) Run(run func(ctx context.Context)) *ContractBackend_SuggestGasPrice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ContractBackend_SuggestGasPrice_Call) Return(_a0 *big.Int, _a1 error) *ContractBackend_SuggestGasPrice_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ContractBackend_SuggestGasPrice_Call) RunAndReturn(run func(context.Context) (*big.Int, error)) *ContractBackend_SuggestGasPrice_Call {
	_c.Call.Return(run)
	return _c
}

// SuggestGasTipCap provides a mock function with given fields: ctx
func (_m *ContractBackend) SuggestGasTipCap(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContractBackend_SuggestGasTipCap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SuggestGasTipCap'
type ContractBackend_SuggestGasTipCap_Call struct {
	*mock.Call
}

// SuggestGasTipCap is a helper method to define mock.On call
//   - ctx context.Context
func (_e *ContractBackend_Expecter) SuggestGasTipCap(ctx interface{}) *ContractBackend_SuggestGasTipCap_Call {
	return &ContractBackend_SuggestGasTipCap_Call{Call: _e.mock.On("SuggestGasTipCap", ctx)}
}

func (_c *ContractBackend_SuggestGasTipCap_Call) Run(run func(ctx context.Context)) *ContractBackend_SuggestGasTipCap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *ContractBackend_SuggestGasTipCap_Call) Return(_a0 *big.Int, _a1 error) *ContractBackend_SuggestGasTipCap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ContractBackend_SuggestGasTipCap_Call) RunAndReturn(run func(context.Context) (*big.Int, error)) *ContractBackend_SuggestGasTipCap_Call {
	_c.Call.Return(run)
	return _c
}

// NewContractBackend creates a new instance of ContractBackend. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewContractBackend(t interface {
	mock.TestingT
	Cleanup(func())
}) *ContractBackend {
	mock := &ContractBackend{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
