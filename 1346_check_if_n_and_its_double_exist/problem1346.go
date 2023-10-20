package problem1346

func checkIfExist(arr []int) bool {
	cache := make(map[int]struct{})

	for _, n := range arr {
		if _, ok := cache[n*2]; ok {
			return true
		}

		if _, ok := cache[n/2]; ok && n%2 == 0 {
			return true
		}

		cache[n] = struct{}{}
	}

	return false
}
