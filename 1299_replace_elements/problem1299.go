package problem1299

// time O(N), space O(1)
func replaceElements(arr []int) []int {
	max := -1
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] > max {
			arr[i], max = max, arr[i]
		} else {
			arr[i] = max
		}
	}
	return arr
}
