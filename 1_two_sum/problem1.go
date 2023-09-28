package problem1

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums)) // num-index

	for i, n := range nums {
		if a, ok := index[target-n]; ok {
			return []int{i, a}
		}

		index[n] = i
	}

	return []int{}
}
