package summary_ranges

import "fmt"

func summaryRanges(nums []int) []string {
	l := len(nums)
	switch l {
	case 0:
		return []string{}
	case 1:
		return []string{fmt.Sprintf("%v", nums[0])}
	}

	res := make([]string, 0)
	for a, b := 0, 1; b <= l; b++ {
		if b == l || nums[b]-nums[b-1] > 1 {
			if b-1 == a {
				res = append(res, fmt.Sprintf("%v", nums[a]))
			} else {
				res = append(res, fmt.Sprintf("%v->%v", nums[a], nums[b-1]))
			}
			a = b
		}
	}
	return res
}
