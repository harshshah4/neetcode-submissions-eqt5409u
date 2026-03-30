/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var checksum func(*TreeNode, int) bool

	checksum = func(root *TreeNode, sum int) bool {
		if root.Left == nil && root.Right == nil {
			if targetSum == sum + root.Val {
				return true
			}
			return false
		} else if root.Right == nil {
			return checksum(root.Left, sum+root.Val)
		} else if root.Left == nil {
			return checksum(root.Right, sum+root.Val)
		} else{
			return checksum(root.Left, sum+root.Val) || checksum(root.Right, sum+root.Val) 
		}
	}
	return checksum(root, 0)
}
