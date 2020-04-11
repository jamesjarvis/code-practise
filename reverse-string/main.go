func reverseString(s []byte) []byte {
	var i int
	var j int = len(s) - 1
	for i < j {
			s[i], s[j] = s[j], s[i]
			i++
			j--
	}
	return s
}
