package problem28

func strStr(haystack string, needle string) int {
	hl, nl := len(haystack), len(needle)

	for i := 0; i < hl-nl+1; i++ {
		if haystack[i:i+nl] == needle {
			return i
		}
	}
	return -1
}
