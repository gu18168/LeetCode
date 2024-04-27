package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + goodNode(root.Left, root.Val) + goodNode(root.Right, root.Val)
}

func goodNode(node *TreeNode, val int) int {
	if node == nil {
		return 0
	}

	isGood := 0
	if node.Val >= val {
		isGood = 1
	}

	val = max(val, node.Val)

	return isGood + goodNode(node.Left, val) + goodNode(node.Right, val)
}
