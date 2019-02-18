package pathSum112

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum, current int) bool {
	if root.Left == nil && root.Right == nil {
		if (root.Val + current) == sum {
			return true
		}
		return false
	}
	current = current+root.Val
	isLeft, isRight := false, false
	if root.Left != nil {
		isLeft = pathSum(root.Left, sum, current)
	}
	if root.Right != nil {
		isRight = pathSum(root.Right, sum, current)
	}
	return isLeft || isRight
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return pathSum(root, sum, 0)
}
