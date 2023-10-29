package problem905

// time O(N), space O(1)
func sortArrayByParity(nums []int) []int {
	e := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[e], nums[i] = nums[i], nums[e]
			e++
		}
	}
	return nums
}
