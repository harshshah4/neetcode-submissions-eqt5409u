/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	_, isbal := isBalLen(root)
	return isbal
}

func isBalLen(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	hL, bL := isBalLen(root.Left)
	hR, bR := isBalLen(root.Right)
	if bL && bR && diff(hL, hR) <= 1 {
		return max(hL, hR) + 1, true
	} 
	return 0, false
}

func diff(hL, hR int) int {
	if hL <= hR {
		return hR - hL
	} 
	return hL - hR
}
