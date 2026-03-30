/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	var array []int
	if root == nil {
		return []int{}
	}
	array = append(array, inorderTraversal(root.Left)...)
	array = append(array, root.Val)
	array = append(array, inorderTraversal(root.Right)...)
	return array
}
