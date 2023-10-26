package problem941

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}

	up := true
	for i := 1; i < len(arr); i++ {
		switch {
		case (up && arr[i] > arr[i-1]) || (!up && arr[i] < arr[i-1]):
			continue
		case up && i > 1 && arr[i] < arr[i-1]:
			up = false
		default:
			return false
		}
	}

	return !up
}
