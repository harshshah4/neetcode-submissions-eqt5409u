func letterCombinations(digits string) []string {
	digitToChar := map[byte]string{
        '2': "abc",
        '3': "def",
        '4': "ghi",
        '5': "jkl",
        '6': "mno",
        '7': "pqrs",
        '8': "tuv",
        '9': "wxyz",
    }

	var out []string
	if len(digits) == 0 {
		return []string{}
	}

	var recurse func(int, []rune) 
	recurse = func(i int, curr []rune) {
		if i == len(digits) {
			fmt.Sprintf("curr %s \n", string(curr))
			out = append(out, strings.Clone(string(curr)))
			return
		}

		for _, c := range digitToChar[digits[i]] {
			curr = append(curr, c)
			recurse(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}
	recurse(0,[]rune{})
	return out
}
