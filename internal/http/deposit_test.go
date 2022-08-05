package http_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arnaz06/deposit"
	handler "github.com/arnaz06/deposit/internal/http"
	"github.com/arnaz06/deposit/mocks"
	"github.com/arnaz06/deposit/testtools"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDepositHandler_get(t *testing.T) {
	tests := []struct {
		testname           string
		mockService        testtools.MockCall
		walletID           string
		expectedStatusCode int
	}{
		{
			testname: "success",
			walletID: "1",
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
			expectedStatusCode: http.StatusOK,
		},
		{
			testname: "error invalid wallerID",
			walletID: "invalid",
			mockService: testtools.MockCall{
				Called: false,
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			testname: "unexpected error",
			walletID: "1",
			mockService: testtools.MockCall{
				Called: true,
				Input: []interface{}{
					mock.Anything,
					int64(1),
				},
				Output: []interface{}{deposit.Deposit{}, errors.New("unexpected error")},
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		mockDepositService := new(mocks.DepositService)
		t.Run(test.testname, func(t *testing.T) {
			if test.mockService.Called {
				mockDepositService.On("Get", test.mockService.Input...).Return(test.mockService.Output...).Once()
			}

			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/deposit/%s", test.walletID), nil)
			rec := httptest.NewRecorder()
			e := echo.New()
			handler.AddDepositHandler(e, mockDepositService)
			e.ServeHTTP(rec, req)

			require.Equal(t, test.expectedStatusCode, rec.Code)
		})
	}
}

func TestDepositHandler_deposit(t *testing.T) {
	dep := deposit.Deposit{
		WalletID:       1,
		Balance:        100,
		AboveThreshold: true,
	}
	tests := []struct {
		testname           string
		mockService        testtools.MockCall
		input              deposit.Deposit
		expectedStatusCode int
	}{
		{
			testname: "success",
			input:    dep,
			mockService: testtools.MockCall{
				Called: true,
				Input:  []interface{}{mock.Anything, dep},
				Output: []interface{}{nil},
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			testname: "unexpected error",
			input:    dep,
			mockService: testtools.MockCall{
				Called: true,
				Input:  []interface{}{mock.Anything, dep},
				Output: []interface{}{errors.New("unexpected error")},
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range tests {
		mockDepositService := new(mocks.DepositService)
		t.Run(test.testname, func(t *testing.T) {
			if test.mockService.Called {
				mockDepositService.On("Deposit", test.mockService.Input...).Return(test.mockService.Output...).Once()
			}

			inputBytes, _ := json.Marshal(test.input)
			req := httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewReader(inputBytes))
			rec := httptest.NewRecorder()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			e := echo.New()
			handler.AddDepositHandler(e, mockDepositService)
			e.ServeHTTP(rec, req)

			require.Equal(t, test.expectedStatusCode, rec.Code)
		})
	}
}
