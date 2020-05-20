func containsDuplicate(nums []int) bool {
	dups := make(map[int]bool)
	for _, v := range nums {
			if _, exists := dups[v]; exists {
					return true
			}
			dups[v] = true
	}
	return false
}