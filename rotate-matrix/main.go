func rotate(matrix [][]int)  {
	l := len(matrix)

	// go through each square
	for i:=0; i<l/2; i++ {
			// go through each side element in this square
			for j:=i; j<l-i-1; j++ {
					// store current cell in temp variable 
					temp := matrix[i][j] 
					
					// move values from left to top 
					matrix[i][j] = matrix[l-1-j][i]
					
					// move values from bottom to left
					matrix[l-1-j][i] = matrix[l-1-i][l-1-j]
					
					// move values from right to bottom
					matrix[l-1-i][l-1-j] = matrix[j][l-1-i]
					
					// move values from temp to right
					matrix[j][l-1-i] = temp
			}   
	}
}