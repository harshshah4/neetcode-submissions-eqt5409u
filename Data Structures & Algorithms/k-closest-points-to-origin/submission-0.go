type pts struct {
	x int
	y int
	dist float64
}
func kClosest(points [][]int, k int) [][]int {
	var p []pts
	var out [][]int
	for _, num := range points {
		p = append(p, pts{
			x: num[0],
			y: num[1],
			dist: math.Sqrt(float64(num[0])*float64(num[0])+float64(num[1])*float64(num[1])),
		})
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].dist < p[j].dist // Use '>' for descending order
	})
	for i, d := range p {
		if i==k {
			break
		}
		out = append(out, []int{d.x, d.y})
	}
	return out
}
