func maxAreaOfIsland(grid [][]int) int {

    var island func([][]int, int, int, int) int
	max := 0
	island = func(grid [][]int, r, c, count int) int {
		if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[0]) || grid[r][c] == 0 {
			return count
		}

		grid[r][c] = 0
		count++

		count = island(grid, r-1, c, count)
		count = island(grid, r+1, c, count)
		count = island(grid, r, c-1, count)
		count = island(grid, r, c+1, count)

		return count
	}

	for i:=0; i<len(grid); i++ {
		for j:=0; j<len(grid[0]); j++ {
			if grid[i][j] == 1 {
				c := island(grid, i, j, 0)
				if max < c {
					max = c
				}
			}
		}
	}
	return max
}
