func isHappy(n int) bool {
	if n == 1 || n == 7 {
			return true
	}
	sum := n
	x := n
	
	for sum > 9 {
			sum = 0
			
			for x > 0 {
					d := x % 10
					sum += d * d
					x /= 10
			}
			
			if sum == 1 {
					return true
			}
			x = sum
	}

	return sum == 7
}
