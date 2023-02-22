package majority_element_2

func majorityElement(nums []int) []int {
	result := make([]int, 0)
	aim := float64(len(nums) / 3)
	stat := make(map[int]float64)
	for _, n := range nums {
		t := stat[n]
		if t <= aim {
			t++
			if t > aim {
				result = append(result, n)
				t++
			}
			stat[n] = t
		}
	}
	return result
}
