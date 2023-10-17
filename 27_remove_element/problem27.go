package problem27

func removeElement(nums []int, val int) int {
	var k int
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}

		if i != k {
			nums[i], nums[k] = nums[k], nums[i]
		}
		k++
	}
	return k
}
