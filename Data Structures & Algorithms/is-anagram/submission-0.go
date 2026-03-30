func isAnagram(s string, t string) bool {
	data := make(map[rune]int)
	if len(s) != len(t) {
		return false
	}
	for _, char := range s {
		data[char]++
	}
	for _, char := range t {
		if _, exists := data[char]; exists {
			data[char]--
			if data[char] == -1 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
