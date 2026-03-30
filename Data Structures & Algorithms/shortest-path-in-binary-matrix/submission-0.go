type pts struct {
	r int
	c int
}
func shortestPathBinaryMatrix(grid [][]int) int {
	queue := []pts{}
	visit := make(map[pts]bool)
	if grid[0][0] != 0 {
		return -1
	}
	queue = append(queue, pts{r: 0, c:0,})
	visit[pts{r:0, c:0,}] = true
	count := 0
	for len(queue) > 0 {
		for _, val := range queue {
			if len(queue) > 1 {
				queue = queue[1:]
			} else {
				queue = []pts{}
			}
			if val.r == len(grid) - 1 && val.c == len(grid) - 1 {
				return count+1
			}
			diff := [][]int{[]int{-1, 0}, []int{1, 0}, []int{0, -1}, []int{0, 1}, []int{-1, -1}, []int{1,1}, []int{-1, 1}, []int{1,-1}}
			for _, each := range diff {
				if val.r+each[0] < 0 || val.c+each[1] < 0 || val.r+each[0] >= len(grid) || val.c+each[1] >= len(grid) || visit[pts{val.r+each[0], val.c+each[1]}] || grid[val.r+each[0]][val.c+each[1]] == 1 {
					continue					
				}
				queue = append(queue, pts{val.r+each[0], val.c+each[1]})
				visit[pts{val.r+each[0], val.c+each[1]}] = true
			}
		}
		count++
	}
	return -1
}
