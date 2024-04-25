package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return depth(root)
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return 1 + max(depth(node.Left), depth(node.Right))
}
