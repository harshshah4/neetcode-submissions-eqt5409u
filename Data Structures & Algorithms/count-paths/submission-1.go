func uniquePaths(m int, n int) int {
	curr := make([]int, n)
	prev := make([]int, n)
    for i:=m-1; i>=0; i-- {
		curr[n-1] = 1
		for j:=n-2; j>=0; j-- {
			curr[j] = prev[j] + curr[j+1]
		}
		prev = curr
	}
	return curr[0]
}
