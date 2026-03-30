/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	var out [][]int

	var ts func(*TreeNode, int)
	ts = func(root *TreeNode, index int) {
		if root == nil {
			return
		}
		if len(out) > index {
			out[index] = append(out[index], root.Val)
		} else {
			newRow := []int{root.Val}
			out = append(out, newRow)
		}
		ts(root.Left, index+1)
		ts(root.Right, index+1)
	}
	ts(root, 0)
	return out
}
