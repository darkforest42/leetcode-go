package invertTree226

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(left)
	return root
}
