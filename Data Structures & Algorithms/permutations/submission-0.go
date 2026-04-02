import "slices"
func permute(nums []int) [][]int {
	out := [][]int{[]int{nums[0]}}
	for i := 1; i< len(nums); i++ {
		temp := [][]int{}
		for j := 0; j < len(out); j++ {
			for k := 0; k <= len(out[j]); k++ {
				t := slices.Insert(out[j], k, nums[i])
				temp = append(temp, t)
			}
		}
		// 1. Create the outer slice with the same length
		out = make([][]int, len(temp))

		// 2. Iterate and copy each inner slice
		for j := range temp {
			out[j] = make([]int, len(temp[j]))
			copy(out[j], temp[j]) // Using built-in copy() for efficiency
		}
	}
	return out
}
