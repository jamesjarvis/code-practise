func rob(nums []int) int {
	if len(nums) == 0 {
			return 0
	}
	
	dp := []int{0,nums[0]}
	for i:= 2; i<=len(nums);i++ {
			dp = append(dp, max(dp[i-1], dp[i-2]+nums[i-1]))
	}
	return dp[len(nums)]
}

func max(a int, b int) int {
	if a > b {
			return a
	}
	return b
}
