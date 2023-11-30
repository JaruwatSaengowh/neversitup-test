package main

import (
	"testing"
)

func TestCountSmileys(t *testing.T) {
	testCases := []struct {
		input    []string
		expected int
	}{
		{[]string{":)", ";(", ";}", ":-D"}, 2},
		{[]string{";D", ":-(", ":-)", ";~)"}, 3},
		{[]string{";]", ":[", ";*", ":$", ";-D"}, 1},
	}

	for _, tc := range testCases {
		result := CountSmilyFace(tc.input)
		if result != tc.expected {
			t.Errorf("For input %v, expected %d but got %d", tc.input, tc.expected, result)
		}
	}
}
