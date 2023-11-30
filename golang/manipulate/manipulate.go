package main

import (
	"fmt"
	"sort"
	"strings"
)

func Manipulate(text []string) []string {
	result := []string{}
	for _, str := range text {
		permute([]rune(str), 0, len(str)-1, &result)
	}
	sort.Strings(result)
	result = removeDuplicates(result)
	return result
}

func permute(text []rune, l, r int, result *[]string) {
	if l == r {
		*result = append(*result, string(text))
	} else {
		for i := l; i <= r; i++ {
			text[l], text[i] = text[i], text[l]
			permute(text, l+1, r, result)
			text[l], text[i] = text[i], text[l]
		}
	}
}

func removeDuplicates(input []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, v := range input {
		if encountered[v] == false {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
}

func main() {
	testCases := []string{"a", "ab", "abc", "aabb"}

	for _, input := range testCases {
		result := Manipulate([]string{input})
		sort.Strings(result)
		expectedResult := fmt.Sprintf("['%s']", strings.Join(result, "', '"))
		fmt.Printf("With input: '%s'\nYour function should return: %s\n", input, expectedResult)
	}
}
