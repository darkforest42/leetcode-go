package balancedBtree

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

func abs(x, y int) int {
	if x > y {
		return x - y
	}else{
		return y - x
	}
}



func maxDepth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	l, leftIsBalance  := maxDepth(root.Left)
	r, rightIsBalance := maxDepth(root.Right)
	if leftIsBalance && rightIsBalance && abs(l, r) < 2 {
		return max(l, r) + 1, true
	}
	return 0, false
}


func isBalanced(root *TreeNode) bool {
	_, isBalanced := maxDepth(root)
	return isBalanced
}
