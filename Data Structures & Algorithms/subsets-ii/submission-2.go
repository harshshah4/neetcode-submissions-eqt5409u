func subsetsWithDup(nums []int) [][]int {

	sort.Ints(nums)
	out := [][]int{}
	var recurse func(int, []int)
	recurse = func(i int, curr []int) {
		if i == len(nums) {
			ccurr := make([]int, len(curr))
			copy(ccurr, curr)
			out = append(out, ccurr)
			return
		}

		curr = append(curr, nums[i])
		recurse(i+1, curr)
		curr = curr[:len(curr)-1]

		var j int
		for j = i; j < len(nums) && nums[j]==nums[i]; j++ {
			continue
		}
		recurse(j, curr)
	}
	recurse(0, []int{})
	return out
}
