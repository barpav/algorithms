package problem1295

import "strconv"

func findNumbers(nums []int) int {
	var result int

	for _, n := range nums {
		if len(strconv.Itoa(n))%2 == 0 {
			result++
		}
	}

	return result
}
