func combine(n int, k int) [][]int {

	out := [][]int{}
	var recurse func(int, []int)
	recurse = func(i int, curr []int) {
		if len(curr) == k {
			ccurr := make([]int, len(curr))
			copy(ccurr,curr)
			out = append(out, ccurr)
			return
		}

		for j := i; j <= n ; j++ {
			curr = append(curr, j)
			recurse(j+1, curr)
			curr = curr[:len(curr)-1]
		}
	}
	recurse(1, []int{})
	return out
}
