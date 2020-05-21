func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		l := 0
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] == needle[j] {
				l += 1
			} else {
				break
			}
		}
		if l == len(needle) {
			return i
		}
	}

	return -1
}