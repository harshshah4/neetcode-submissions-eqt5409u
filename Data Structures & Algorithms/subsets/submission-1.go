func subsets(nums []int) [][]int {
	var out [][]int
	var set func(int)
	set = func(i int) {
		if i == len(nums) - 1 {
			out = [][]int{[]int{}, []int{nums[i]}}
			return
		}
		set(i+1)
		var arr [][]int
		for _, val := range out {
			arr = append(arr, val)
			cp := make([]int, len(val))
			copy(cp, val)
			cp = append(cp, nums[i])
			arr = append(arr, cp)
		}
		out = arr
		return
	}
	set(0)
	return out
}
