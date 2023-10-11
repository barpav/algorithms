package problem485

func findMaxConsecutiveOnes(nums []int) int {
	var cur, max int

	for _, n := range nums {
		cur += n

		if n == 0 {
			if cur > max {
				max = cur
			}

			cur = 0
		}
	}

	if cur > max {
		max = cur
	}

	return max
}
