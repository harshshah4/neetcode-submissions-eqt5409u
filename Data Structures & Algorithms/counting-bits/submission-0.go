func countBits(n int) []int {
	out := make([]int, n+1)
	pow := 1
	out[0]=0
	for i:=1; i<=n; i++ {
		if i == 1 {
			out[i] = 1
		}
		if i == 2 * pow {
			pow = i
		}
		if i % 2 == 0 {
			out[i] = out[i/2]
			continue
		}
		out[i] = out[i-pow] + out[pow]
	}
	return out
}
