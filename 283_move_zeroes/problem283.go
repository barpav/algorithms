package problem283

// time O(N), space O(1)
func moveZeroes(nums []int) {
	v := -1
	for i := 0; i < len(nums); i++ {
		if v == -1 && nums[i] == 0 {
			v = i
			continue
		}

		if nums[i] != 0 && v != -1 {
			nums[v], nums[i] = nums[i], nums[v]
			v++
		}
	}
}
