func getConcatenation(nums []int) []int {
	n := len(nums)
	numsN := make([]int, n*2)
    for i, num := range nums {
		numsN[i] = num
		numsN[i+n] = num
	}
	return numsN
}
