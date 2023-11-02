package problem9

import "strconv"

func isPalindrome(x int) bool {
	v := strconv.Itoa(x)
	l, r := 0, len(v)-1
	for l < r {
		if v[l] != v[r] {
			return false
		}
		l++
		r--
	}
	return true
}
