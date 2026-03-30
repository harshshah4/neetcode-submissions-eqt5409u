/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
    if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			maxNode := findMax(root.Left)
			root.Val = maxNode.Val
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	return root
}

func findMax(root *TreeNode) *TreeNode {
	for root != nil && root.Right != nil {
		root = root.Right
	}
	return root
}
