package main

import (
	"testing"
)

func TestFindOdd(t *testing.T) {
	testCases := []struct {
		input    []int
		expected int
	}{
		{[]int{7}, 7},
		{[]int{0}, 0},
		{[]int{1, 1, 2}, 2},
		{[]int{0, 1, 0, 1, 0}, 0},
		{[]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}, 4},
	}

	for _, tc := range testCases {
		result := FindOddNumber(tc.input)
		if result != tc.expected {
			t.Errorf("For input %v, expected %d but got %d", tc.input, tc.expected, result)
		}
	}
}
