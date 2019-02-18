package btreeDiameter543

//Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxDepth(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	l, leftDiameter := maxDepth(root.Left)
	r, rightDiameter := maxDepth(root.Right)
	return max(l, r) + 1, max(l+r, max(leftDiameter, rightDiameter))
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, diameter := maxDepth(root)
	return diameter
}