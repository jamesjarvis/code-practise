def rotateMatrix(mat): 
    n = len(mat)
    for x in range(0, int(n / 2)):
        for y in range(x, n-x-1): 
              
            # store current cell in temp variable 
            temp = mat[x][y] 
  
            # move values from right to top 
            mat[x][y] = mat[y][n-1-x] 
  
            # move values from bottom to right 
            mat[y][n-1-x] = mat[n-1-x][n-1-y] 
  
            # move values from left to bottom 
            mat[n-1-x][n-1-y] = mat[n-1-y][x] 
  
            # assign temp to left 
            mat[n - 1 - y][x] = temp
            
    return mat