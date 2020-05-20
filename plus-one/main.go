func plusOne(digits []int) []int {
	//     work backwards dummy
			for i:= len(digits)-1; i >= 0; i -= 1 {
					if digits[i] == 9 {
							digits[i] = 0
					} else {
							digits[i] += 1
							return digits
					}
			}
			digits = append([]int{1}, digits...)
			return digits
	}