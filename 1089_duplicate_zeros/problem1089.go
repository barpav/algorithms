package problem1089

func duplicateZeros(arr []int) {
	i, l := 0, len(arr)
	for i < l {
		if arr[i] != 0 {
			i++
			continue
		}
		arr = append(arr[:i+1], arr[i:l-1]...)
		i += 2
	}
}
