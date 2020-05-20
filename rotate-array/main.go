func rotate(nums []int, k int)  {
	n := len(nums)
	k = k % n
	balls := append(nums[n-k:], nums[:n-k]...)
	for i, v := range balls {
			nums[i] = v
	}
}