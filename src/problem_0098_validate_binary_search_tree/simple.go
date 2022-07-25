package solution

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt32-1, math.MaxInt32+1)
}

func isValid(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}

	if int64(node.Val) >= max || int64(node.Val) <= min {
		return false
	}

	return isValid(node.Left, min, int64(node.Val)) && isValid(node.Right, int64(node.Val), max)
}
