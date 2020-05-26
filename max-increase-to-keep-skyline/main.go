func maxIncreaseKeepingSkyline(grid [][]int) int {
	var sum int
	var xmax, ymax []int
	n := len(grid)
	for i:=0; i<n; i++ {
			// initialise xmax[i] with initial value
			xmax = append(xmax, grid[i][0])
			// initialise ymax[i] with initial value
			ymax = append(ymax, grid[0][i])
			
			for j:=1; j<n; j++ {
					xmax[i] = max(xmax[i], grid[i][j])
					
					ymax[i] = max(ymax[i], grid[j][i])
			}
	}
	
	// repopulate
	for i:=0; i<n; i++ {        
			for j:=0; j<n; j++ {
					sum += min(xmax[i], ymax[j]) - grid[i][j]
			}
	}
	
	return sum
}

func max(a int, b int) int {
	if a > b {
			return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
			return a
	}
	return b
}