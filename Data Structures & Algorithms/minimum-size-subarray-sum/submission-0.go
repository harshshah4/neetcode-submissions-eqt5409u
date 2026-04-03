func minSubArrayLen(target int, nums []int) int {
	length := len(nums)+1
	j:=0
	sum := 0
	for i, val := range nums {
		sum += val
		for sum >= target {
			if i-j+1 < length  {
				length = i-j+1
			}
			sum -= nums[j]
			j++
		}
	}
	if len(nums)+1 == length {
		return 0
	}
	return length
}
