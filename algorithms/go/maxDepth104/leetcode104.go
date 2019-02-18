package maxDepth104

//Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return min(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
