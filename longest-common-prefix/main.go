func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
			return ""
	}
	var i int
	minlength := len(strs[0])
	for l := 1; l < len(strs); l++ {
			if len(strs[l]) < minlength {
					minlength = len(strs[l])
			}
	}
	for i < minlength {
			tempChar := strs[0][i]
			for a := 0; a < len(strs); a++ {
					if strs[a][i] != tempChar {
							return strs[0][:max(0, i)]
					}
			}
			i++
	}
	return strs[0][:minlength]
}

func max(a int, b int) int {
	if a > b {
			return a
	}
	return b
}