func romanToInt(s string) int {
	var sum int
	for i := 0; i < len(s); i++ {
			if i < len(s)-1 && value(s[i]) < value(s[i+1]) {
					sum -= value(s[i])
			} else {
					sum += value(s[i])
			}
	}
	return sum
}

func value(c byte) int {
	if c == 'I' {
			return 1
	}
	if c == 'V' {
			return 5
	}
	if c == 'X' {
			return 10
	}
	if c == 'L' {
			return 50
	}
	if c == 'C' {
			return 100
	}
	if c == 'D' {
			return 500
	}
	if c == 'M' {
			return 1000
	}
	return 0
}