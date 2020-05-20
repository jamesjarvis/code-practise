func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
			return 0
	}
	
	m := len(grid)
	n := len(grid[0])
	
	// initialise 2d slice
	dp := make([][]int, m) 
	for nn := 0; nn < m; nn++ {
			dp[nn] = make([]int, n)
	}
	
	dp[0][0] = grid[0][0]
	
	// initialise top row
	for i := 1; i < n; i++ {
			dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	
	// initialise left column
	for j := 1; j < m; j++ {
			dp[j][0] = dp[j-1][0] + grid[j][0]
	}
	
	// fill up the dp table
	for i := 1; i < m; i++ {
			for j := 1; j < n; j++ {
					if dp[i-1][j] > dp[i][j-1] {
							dp[i][j] = dp[i][j-1] + grid[i][j]
					} else {
							dp[i][j] = dp[i-1][j] +  grid[i][j]
					}
			}
	}
	
	return dp[m-1][n-1]
}
