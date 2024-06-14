package fizzbuzz

import (
	"testing"
)

// FizzBuzz関数のテストケース
func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{1, "1"},
		{3, "Fizz"},
		{5, "Buzz"},
		{15, "FizzBuzz"},
		{30, "FizzBuzz"},
		{7, "7"},
	}

	for _, tc := range testCases {
		result := FizzBuzz(tc.input)
		if result != tc.expected {
			t.Errorf("Expected %s, but got %s for input %d",
				tc.expected, result, tc.input)
		}
	}
}
