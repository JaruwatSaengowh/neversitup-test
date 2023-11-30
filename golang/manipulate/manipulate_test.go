package main

import (
	"reflect"
	"testing"
)

func TestManipulate(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"abc", "acb", "bac", "bca", "cab", "cba"}},
		{"aabb", []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}},
	}

	for _, test := range tests {
		result := Manipulate([]string{test.input})
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %s, expected %v, but got %v", test.input, test.expected, result)
		}
	}
}
