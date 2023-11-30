package main

import (
	"fmt"
)

func FindOddNumber(text []int) int  {
	frequency := make(map[int]int)

    for _, num := range text {
		frequency[num]++
	}

	for key, value := range frequency {
		if value%2 != 0 {
			return key
		}
	}

	return 0
}

func main() {
	testCases := [][]int{
		{7},
		{0},
		{1, 1, 2},
		{0, 1, 0, 1, 0},
		{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
	}

	for _, testCase := range testCases {
		expectedResult := FindOddNumber(testCase)
		fmt.Printf("%v should return %d, because it occurs %d time(s) (which is odd).\n", testCase, expectedResult, frequency(testCase, expectedResult))
	}
}

func frequency(arr []int, num int) int {
	count := 0
	for _, value := range arr {
		if value == num {
			count++
		}
	}
	return count
}
