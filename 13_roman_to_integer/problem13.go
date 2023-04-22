package problem13

func romanToInt(s string) int {
	symbols := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sl, result := len(s), 0
	for i := 0; i < sl; i++ {
		if i+1 < sl && symbols[s[i+1]] > symbols[s[i]] {
			result -= symbols[s[i]]
		} else {
			result += symbols[s[i]]
		}
	}

	return result
}
