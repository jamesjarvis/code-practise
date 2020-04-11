func twoSum(nums []int, target int) []int {
	mapped := make(map[int]int)
	for i, v := range nums {
			complement := target - v
			k, ok := mapped[complement]
			if ok {
					return []int{k,i}
			}
			mapped[v] = i
	}
	return []int{}
}