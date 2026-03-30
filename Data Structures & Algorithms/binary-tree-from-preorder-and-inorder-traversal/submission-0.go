/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	var imap map[int]int
	imap = make(map[int]int)
	for value, key := range inorder {
		imap[key] = value
	}
	i := 0
	var formTree func(int, int) *TreeNode
	formTree = func(start, end int) *TreeNode {
		if start > end || i >= len(preorder) {
			return nil
		}
		root := &TreeNode{
			Val: preorder[i],
		}
		curr := preorder[i]
		i++
		root.Left = formTree(start, imap[curr]-1)
		root.Right = formTree(imap[curr]+1, end)
		return root
	}
	return formTree(0, len(preorder)-1)
}
