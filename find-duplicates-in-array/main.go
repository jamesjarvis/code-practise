func findDuplicates(nums []int) []int {
	var singles []int
	for i:=0; i<len(nums);i++ {
			temp := abs(nums[i]) - 1
			if nums[temp] < 0 {
					singles = append(singles, temp+1)
			} else {
					nums[temp] = 0 - nums[temp]
			}
	}
	return singles
}

func abs(a int) int {
	if a > 0 {
			return a
	}
	return 0 - a
}
