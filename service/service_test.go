package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/arnaz06/deposit"
	"github.com/arnaz06/deposit/mocks"
	"github.com/arnaz06/deposit/pb"
	"github.com/arnaz06/deposit/service"
	"github.com/arnaz06/deposit/testtools"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDepositService_deposit(t *testing.T) {
	tests := []struct {
		testname      string
		input         deposit.Deposit
		mockService   testtools.MockCall
		expectedError error
	}{
		{
			testname: "success",
			input: deposit.Deposit{
				WalletID: 1,
				Balance:  100,
			},
			mockService: testtools.MockCall{
				Called: true,
				Input: []interface{}{
					mock.Anything,
					&pb.Deposit{
						WalletId: 1,
						Balance:  100,
					},
				},
				Output: []interface{}{nil},
			},
			expectedError: nil,
		},
		{
			testname: "unexpected error from repo",
			input: deposit.Deposit{
				WalletID: 1,
				Balance:  100,
			},
			mockService: testtools.MockCall{
				Called: true,
				Input: []interface{}{
					mock.Anything,
					&pb.Deposit{
						WalletId: 1,
						Balance:  100,
					},
				},
				Output: []interface{}{errors.New("unexpected error")},
			},
			expectedError: errors.New("unexpected error"),
		},
	}
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			mockDepositService := new(mocks.DepositRepository)
			if test.mockService.Called {
				mockDepositService.On("Deposit", test.mockService.Input...).Return(test.mockService.Output...).Once()
			}

			depositService := service.NewDepositService(mockDepositService)
			err := depositService.Deposit(context.Background(), test.input)
			if test.expectedError != nil {
				require.EqualError(t, err, test.expectedError.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestDepositService_get(t *testing.T) {
	tests := []struct {
		testname       string
		input          int64
		mockService    testtools.MockCall
		expectedResult deposit.Deposit
		expectedError  error
	}{
		{
			testname: "success",
			input:    1,
			mockService: testtools.MockCall{
				Called: true,
				Input: []interface{}{
					mock.Anything,
					int64(1),
				},
				Output: []interface{}{deposit.Deposit{
					WalletID:       1,
					Balance:        100,
					AboveThreshold: true,
				}, nil},
			},
			expectedResult: deposit.Deposit{
				WalletID:       1,
				Balance:        100,
				AboveThreshold: true,
			},
			expectedError: nil,
		},
		{
			testname: "unexpected error from repo",
			input:    1,
			mockService: testtools.MockCall{
				Called: true,
				Input: []interface{}{
					mock.Anything,
					int64(1),
				},
				Output: []interface{}{deposit.Deposit{}, errors.New("unexpected error")},
			},
			expectedResult: deposit.Deposit{},
			expectedError:  errors.New("unexpected error"),
		},
	}
	for _, test := range tests {
		t.Run(test.testname, func(t *testing.T) {
			mockDepositService := new(mocks.DepositRepository)
			if test.mockService.Called {
				mockDepositService.On("Get", test.mockService.Input...).Return(test.mockService.Output...).Once()
			}

			depositService := service.NewDepositService(mockDepositService)
			res, err := depositService.Get(context.Background(), test.input)
			if test.expectedError != nil {
				require.EqualError(t, err, test.expectedError.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.expectedResult, res)
		})
	}
}
