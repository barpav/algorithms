package problem977

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	l, r, lv, rv := 0, len(nums)-1, 0, 0

	for i := r; i >= 0; i-- {
		lv, rv = nums[l]*nums[l], nums[r]*nums[r]

		if rv > lv {
			result[i], r = rv, r-1
		} else {
			result[i], l = lv, l+1
		}
	}

	return result
}
