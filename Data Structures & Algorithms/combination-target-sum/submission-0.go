func combinationSum(nums []int, target int) [][]int {
	var out [][]int
    
	var repeat func(int, []int, int)
	repeat = func(start int, res []int, sum int) {
		if sum == target {
			out = append(out, res)
			return
		}
		if sum > target {
			return
		}
		for i:= start; i < len(nums); i++ {
			cp := make([]int, len(res))
		    copy(cp, res)
			repeat(i, append(cp, nums[i]), sum+nums[i])
		}
		return
	}

	for i, _ := range nums {
		repeat(i, []int{nums[i]}, nums[i])
	}
	return out
}
