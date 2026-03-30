func twoSum(nums []int, target int) []int {
	data := make(map[int]int)
    for i, num := range nums {
		if val, exists := data[target-num]; exists {
			return []int{val, i}
		}
		data[num] = i
	}
	return []int{}
}
