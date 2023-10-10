package problem3

func lengthOfLongestSubstring(s string) int {
	var head, l, max int
	idx := make(map[rune]int, len(s))

	for i, ch := range s {
		if ii, ok := idx[ch]; ok && head < ii+1 {
			head = ii + 1
		}

		l = i - head + 1

		if l > max {
			max = l
		}

		idx[ch] = i
	}

	return max
}
