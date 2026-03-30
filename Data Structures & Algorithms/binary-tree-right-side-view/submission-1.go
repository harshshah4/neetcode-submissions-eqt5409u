/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
	out := []int{}
	if root == nil {
		return out
	}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		out = append(out, queue[len(queue)-1].Val)
		var newQueue []*TreeNode
		for _, i := range queue {
			if i.Left != nil {
				newQueue = append(newQueue, i.Left)
			}
			if i.Right != nil {
				newQueue = append(newQueue, i.Right)
			}
		}
		queue = newQueue[:]
	}
	return out
}
