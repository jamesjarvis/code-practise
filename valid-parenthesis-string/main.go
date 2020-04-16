func checkValidString(s string) bool {
	// This cannot be solved easily with the classic stack approach for ()
	var lo int
	var hi int
	for _, v := range s {
			if v == '(' {
					lo ++
			} else {
					lo --
			}
			if v != ')' {
					hi ++
			} else {
					hi --
			}
			if hi < 0 {
					break
			}
			lo = max(0, lo)
	}
	return lo == 0
}

func max(a int, b int) int {
	if a > b {
			return a
	}
	return b
}