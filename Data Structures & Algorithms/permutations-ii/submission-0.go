import "slices"
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	out := [][]int{[]int{nums[0]}}
	for i := 1; i< len(nums); i++ {
		temp := [][]int{}
		for j := 0; j < len(out); j++ {
			for k := 0; k <= len(out[j]); k++ {
				t := slices.Insert(out[j], k, nums[i])
				temp = append(temp, t)
				// fmt.Printf("j %d k %d i %d \t", j ,k , i)
				// fmt.Printf("out[%d][%d] %d\n",j,k, out[j][k])
				if k != len(out[j]) && out[j][k] == nums[i] {
					break
				}
			}
		}
		out = make([][]int, len(temp))
		for j := range temp {
			out[j] = make([]int, len(temp[j]))
			copy(out[j], temp[j])
		}
	}
	return out
}
