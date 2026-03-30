/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    i := 0
	val := -1
	var ksm func(*TreeNode)
	ksm = func(root *TreeNode) {
		if root == nil {
			return
		}
		ksm(root.Left)
		i++
		if i == k {
			val = root.Val
		}
		ksm(root.Right)
	}
	ksm(root)
	return val
}
