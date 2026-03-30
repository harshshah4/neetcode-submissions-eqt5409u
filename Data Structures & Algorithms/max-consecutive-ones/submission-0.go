func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	count := 0
	for i:= 0; i< len(nums); i++ {
		if nums[i] != 1 {
			if max < count {
				max = count
			}
			count = 0
			continue
		}
		count++
	}
	if max < count {
		return count
	}
	return max
}
