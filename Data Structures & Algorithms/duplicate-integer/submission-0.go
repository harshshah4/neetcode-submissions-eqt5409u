func hasDuplicate(nums []int) bool {
    data := make(map[int]struct{})

	for _, num := range nums {
		if _, exists := data[num]; exists {
			return true
		}
		data[num] = struct{}{}
	}
	return false
}
