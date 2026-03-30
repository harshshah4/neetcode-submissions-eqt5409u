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
	arr := inorderTraversal(root.Left)
	array = append(array, arr...)
	array = append(array, root.Val)
	arr = inorderTraversal(root.Right)
	array = append(array, arr...)
	return array
}
