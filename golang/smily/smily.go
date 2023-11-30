package main

import (
	"fmt"
	"regexp"
	"strings"
)

func CountSmilyFace(text []string) int {
	count := 0
	re := regexp.MustCompile(`[:;][-~]?[)D]`)

	for _, face := range text {
		if re.MatchString(face) {
			count++
		}
	}

	return count
}

func formatTestCase(testCase []string) string {
	return fmt.Sprintf("['%s']", strings.Join(testCase, "', '"))
}

func main() {
	testCases := [][]string{
		{":)", ";(", ";}", ":-D"},
		{";D", ":-(", ":-)", ";~)"},
		{";]", ":[", ";*", ":$", ";-D"},
	}

	for _, testCase := range testCases {
		expectedResult := CountSmilyFace(testCase)
		fmt.Printf("countSmileys(%v) // should return %d;\n", formatTestCase(testCase), expectedResult)
	}
}

