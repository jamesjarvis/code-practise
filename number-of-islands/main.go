func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
			return 0
	}
	rowlen, collen := len(grid), len(grid[0])
	// visited array
	visited := make([][]bool, rowlen)
	for i := range visited {
			visited[i] = make([]bool, collen)
	}
	var count int
	for i:=0; i<rowlen; i++ {
			for j:=0; j<collen; j++ {
					// if island cell not visited yet, then visit all cells in the island and increment count
					if grid[i][j] == '1' && !visited[i][j] {
							visited = dfs(grid, i, j, visited, rowlen, collen)
							count++
					}
			}
	}
	return count
}

func dfs(grid [][]byte, i int, j int, visited [][]bool, rowlen int, collen int) [][]bool {
	visited[i][j] = true
	rows := []int{-1,0,0,1}
	cols := []int{0,-1,1,0}
	for k:=0; k<len(rows); k++ {
			// Check that the new index is valid and not visited before
			if i+rows[k]>=0 && i+rows[k]<rowlen && j+cols[k]>=0 && j+cols[k]<collen && !visited[i+rows[k]][j+cols[k]] && grid[i+rows[k]][j+cols[k]] == '1' {
					visited = dfs(grid, i+rows[k], j+cols[k], visited, rowlen, collen)
			}
	}
	return visited
}