func twoSum(nums []int, target int) []int {
	for a := 0; a < len(nums)-1; a++ {
			for b := a+1; b<len(nums); b++ {
					if nums[a] + nums[b] == target {
							return []int{a,b}
					}
			}
	}
	return []int{}
}