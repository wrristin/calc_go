package calculation_test

import (
	"calc_service/internal/calculation"
	"testing"
)

func TestCalc(t *testing.T) {
	testCases := []struct {
		expression     string
		expectedResult string
		expectError    bool
	}{
		{
			expression:     "3 + 5 * (2 - 8)",
			expectedResult: "-13.000000",
			expectError:    false,
		},
		{
			expression:     "10 / 2",
			expectedResult: "5.000000",
			expectError:    false,
		},
		{
			expression:     "3 + a",
			expectedResult: "",
			expectError:    true,
		},
		{
			expression:     "3 / 0",
			expectedResult: "",
			expectError:    true,
		},
		{
			expression:     "(3 + 4",
			expectedResult: "",
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		result, err := calculation.Calc(tc.expression)
		if (err != nil) != tc.expectError {
			t.Fatalf("For expression %s, expected error: %v, got: %v", tc.expression, tc.expectError, err)
		}
		if !tc.expectError && result != tc.expectedResult {
			t.Fatalf("For expression %s, expected result: %s, got: %s", tc.expression, tc.expectedResult, result)
		}
	}
}
