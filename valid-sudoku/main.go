func isValidSudoku(board [][]byte) bool {
	//  naive approach, but apparently the fastest?
			
			var blocks []map[byte]bool
			
			for x := 0; x < 9; x += 1 {
					xmap := make(map[byte]bool)
					ymap := make(map[byte]bool)
					if x % 3 == 0 {
							blocks = make([]map[byte]bool, 3)
							for ix := 0; ix < 3; ix += 1 {
									blocks[ix] = make(map[byte]bool)
							}
					}
	
					for y := 0; y < 9; y += 1 {
	//          x value check
							if board[x][y] != '.' {
									//  x direction check
									if _, foundx := xmap[board[x][y]]; foundx {
											return false
									}
									xmap[board[x][y]] = true
	
									//  block check
									if _, found := blocks[y/3][board[x][y]]; found {
											return false
									}
									blocks[y/3][board[x][y]] = true
							}
					
							
	//          y value check
							if board[y][x] != '.' {
									//  y direction check
									if _, foundy := ymap[board[y][x]]; foundy {
											return false
									}
									ymap[board[y][x]] = true   
							}
					}
			}
			return true
	}