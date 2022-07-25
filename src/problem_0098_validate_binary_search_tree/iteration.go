package solution

import "math"

// Incremental middle-order traversal
func isValidBSTIter(root *TreeNode) bool {
	stack := make([]*TreeNode, 0)

	last := math.MinInt32 - 1

	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		l := len(stack)
		root = stack[l-1]
		stack = stack[:l-1]

		if root.Val <= last {
			return false
		}
		last = root.Val
		root = root.Right
	}

	return true
}
