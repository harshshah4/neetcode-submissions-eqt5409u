func numIslands(grid [][]byte) int {
    
	var island func([][]byte, int, int)
	count := 0
	island = func(grid [][]byte, r, c int) {
		if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'

		island(grid, r-1, c)
		island(grid, r+1, c)
		island(grid, r, c-1)
		island(grid, r, c+1)

		return
	}

	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[0]); j++ {
			if grid[i][j] == '1' {
				island(grid, i, j)
				count++
			}
		}
	}
	return count
}
