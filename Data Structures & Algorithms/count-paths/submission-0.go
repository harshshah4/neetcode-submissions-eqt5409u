func uniquePaths(m int, n int) int {
	grid := make([][]int, m) // Create the outer slice
    for i := range grid {
        grid[i] = make([]int, n) // Initialize each inner slice (row)
    }
	prev := make([]int, n)
    for i:=m-1; i>=0; i-- {
		grid[i][n-1] = 1
		for j:=n-2; j>=0; j-- {
			grid[i][j] = prev[j] + grid[i][j+1]
		}
		prev = grid[i]
	}
	return grid[0][0]
}
