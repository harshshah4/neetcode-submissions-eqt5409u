type SegmentTree struct {
	sum int
	left *SegmentTree
	right *SegmentTree
	start int
	end int
}

func NewSegmentTree(nums []int) *SegmentTree {
	var createTree func(int, int) *SegmentTree 
	createTree = func(l, r int) *SegmentTree {
		if l == r {
			return &SegmentTree {
				sum: nums[l],
				start: l,
				end: r,
			}
		}
		m := (l + r) / 2
		left := createTree(l,m)
		right := createTree(m+1, r)
		return &SegmentTree {
			sum : left.sum + right.sum,
			left: left,
			right: right,
			start: l,
			end: r,
		}
	}
	return createTree(0, len(nums)-1)
}

func (st *SegmentTree) Update(index, val int) {
	var updateTree func(*SegmentTree)
	updateTree = func(curr *SegmentTree) {
		if curr.start == index && curr.end == index {
			curr.sum = val
			return
		}
		if curr.end < index || curr.start > index {
			return
		}
		updateTree(curr.left)
		updateTree(curr.right)
		curr.sum = curr.left.sum + curr.right.sum
		return
	}
	updateTree(st)
	return
}

func (st *SegmentTree) Query(l, r int) int {
	var queryTree func(int, int, *SegmentTree) int
	sum := 0
	queryTree = func(s, e int, curr *SegmentTree) int {
		if s == curr.start && curr.end == e {
			return curr.sum
		}
		m := (curr.start + curr.end) / 2
		if m + 1 > e {
			sum = queryTree(max(curr.start, s), e, curr.left)
		} else if m < s {
			sum = queryTree(s, min(curr.end, e), curr.right)
		} else {
			sum = queryTree(max(curr.start, s), m, curr.left) + queryTree(m+1,min(curr.end,e),curr.right)
		}
		return sum
	}
	return queryTree(l, r, st)
}

