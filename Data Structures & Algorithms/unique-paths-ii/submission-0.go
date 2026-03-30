type pts struct {
	r int
	c int
}
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	matrix := make([][]int, len(obstacleGrid))
    for i := range matrix {
        matrix[i] = make([]int, len(obstacleGrid[0]))
    }
	cache := make(map[pts]int)
	var traverse func(int, int) int
	traverse = func(r,c int) int {
		if r < 0 || c < 0 || r>=len(matrix) || c >= len(matrix[0]) {
			return 0
		}
		if obstacleGrid[r][c] == 1 {
			return 0
		}
		if r == len(matrix) - 1 && c == len(matrix[0]) - 1 {
			return 1
		}

		if val, exists := cache[pts{r,c}]; exists {
			return val
		}
		cache[pts{r,c}] = traverse(r+1, c) + traverse(r, c+1)
		return cache[pts{r,c}]
	}
	return traverse(0,0)
}
