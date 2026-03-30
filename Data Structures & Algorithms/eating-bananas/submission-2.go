import "slices"
func minEatingSpeed(piles []int, h int) int {
	end := slices.Max(piles)
	start := 1
	if len(piles) > h {
		return -1
	}
	lastValid := end
	for start <= end {
		mid := (start + end) /2
		if validate(piles, mid, h) {
			end = mid - 1
			lastValid = mid
		} else {
			start = mid + 1
		}
	}
	return lastValid
}

func validate(piles []int, eat, threshold int) bool {
	count := 0
	for _, pile := range piles {
		if pile % eat == 0 {
			count += pile/eat
		} else {
			count += pile/eat + 1
		}
	}
	if count <= threshold {
		return true
	}
	return false
}
