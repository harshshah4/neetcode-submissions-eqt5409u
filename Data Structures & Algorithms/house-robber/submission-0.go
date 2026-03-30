func rob(nums []int) int {

	cache := make(map[int]int)
    
	var choose func(int) int

	choose = func(i int) int {
		if i >= len(nums) || i < 0 {
			return 0
		}
		if val, exists := cache[i]; exists {
			return val
		}
		cache[i] = max(choose(i+2), choose(i+3)) + nums[i]
		return cache[i]
	}
	return max(choose(0), choose(1))
}
