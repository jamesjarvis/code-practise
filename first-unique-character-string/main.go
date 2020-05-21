func firstUniqChar(s string) int {
	// build frequency map, using an array of only lowercase letter values
	m := make([]uint16, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a'] += 1
	}
	// get index
	for j := 0; j < len(s); j++ {
		if m[s[j]-'a'] == 1 {
			return j
		}
	}
	return -1
}