// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	deposit "github.com/arnaz06/deposit"
	mock "github.com/stretchr/testify/mock"
)

// DepositService is an autogenerated mock type for the DepositService type
type DepositService struct {
	mock.Mock
}

// Deposit provides a mock function with given fields: ctx, input
func (_m *DepositService) Deposit(ctx context.Context, input deposit.Deposit) (deposit.Deposit, error) {
	ret := _m.Called(ctx, input)

	var r0 deposit.Deposit
	if rf, ok := ret.Get(0).(func(context.Context, deposit.Deposit) deposit.Deposit); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(deposit.Deposit)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, deposit.Deposit) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, walletID
func (_m *DepositService) Get(ctx context.Context, walletID int64) (deposit.Deposit, error) {
	ret := _m.Called(ctx, walletID)

	var r0 deposit.Deposit
	if rf, ok := ret.Get(0).(func(context.Context, int64) deposit.Deposit); ok {
		r0 = rf(ctx, walletID)
	} else {
		r0 = ret.Get(0).(deposit.Deposit)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, walletID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
