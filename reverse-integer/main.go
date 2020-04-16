func reverse(x int) int {
	MaxInt32 := 1<<31 - 1
	MinInt32 := -1 << 31
	var y int
	
	for x != 0 {
			y = y*10 + x%10
			if y > MaxInt32 || y < MinInt32 {
					return 0
			}
			x /= 10
	}

	return y
}