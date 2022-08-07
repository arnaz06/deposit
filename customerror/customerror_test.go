package customerror_test

import (
	"testing"

	"github.com/arnaz06/deposit/customerror"
	"github.com/stretchr/testify/require"
)

func TestErrorNotFound(t *testing.T) {
	tests := []struct {
		testName      string
		inputedError  error
		expectedError string
	}{
		{
			testName:      "success",
			inputedError:  customerror.ErrorNotFoundf("formated error %s", "true"),
			expectedError: "formated error true",
		},
		{
			testName:      "success with multiple args",
			inputedError:  customerror.ErrorNotFoundf("formated error %s %d", "true", 1),
			expectedError: "formated error true 1",
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			errString, ok := test.inputedError.(customerror.ErrorNotFound)
			if ok {
				require.Equal(t, test.expectedError, errString.Error())
				return
			}
		})
	}

}

func TestValidationError(t *testing.T) {
	tests := []struct {
		testName      string
		inputedError  error
		expectedError string
	}{
		{
			testName:      "success",
			inputedError:  customerror.ValidationErrorf("formated error %s", "true"),
			expectedError: "formated error true",
		},
		{
			testName:      "success with multiple args",
			inputedError:  customerror.ValidationErrorf("formated error %s %d", "true", 1),
			expectedError: "formated error true 1",
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			errString, ok := test.inputedError.(customerror.ValidationError)
			if ok {
				require.Equal(t, test.expectedError, errString.Error())
				return
			}
		})
	}

}
