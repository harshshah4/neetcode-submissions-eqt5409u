func removeElement(nums []int, val int) int {
	goback := 0
    for i, num := range nums {
		if num == val {
			goback++
			continue
		}
		nums[i-goback] = num
	}
	return len(nums)-goback
}
