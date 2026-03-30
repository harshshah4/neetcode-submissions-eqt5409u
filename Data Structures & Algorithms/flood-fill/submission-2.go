func floodFill(image [][]int, sr int, sc int, color int) [][]int {

	type rc struct {
		r int
		c int
	}

	match := image[sr][sc]
	if match == color {
		return image
	}
	var change func([][]int, int, int)

	change = func(image [][]int, r, c int) {
		if r < 0 || c < 0 || r >= len(image) || c >= len(image[0]) || image[r][c] != match {
			return
		}

		image[r][c] = color
		change(image, r+1, c)
		change(image, r-1, c)
		change(image, r, c+1)
		change(image, r, c-1)

		return
	}

	change(image, sr, sc)
	return image
}
