func missingNumber(nums []int) int {
	// my simple idea is you can just calculate the sum of all numbers below len+1, then subtract each element from this and that should give you the missing number?
	l := len(nums)
	s := (l*(l+1))/2
	for i := 0; i < l; i++ {
			s -= nums[i]
	}

	return s
}
