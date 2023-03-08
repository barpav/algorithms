package search_insert_position

func searchInsert(nums []int, target int) int {
	l, r, m := 0, len(nums)-1, 0
	for l <= r {
		m = (l + r) / 2
		switch {
		case target < nums[m]:
			r = m - 1
		case target > nums[m]:
			l = m + 1
		default:
			return m
		}
	}
	return l
}
