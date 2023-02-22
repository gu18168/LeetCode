package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	node := new(TreeNode)

	lVal, rVal := 0, 0

	var ll, lr, rl, rr *TreeNode

	if root1 != nil {
		lVal = root1.Val
		ll, lr = root1.Left, root1.Right
	}

	if root2 != nil {
		rVal = root2.Val
		rl, rr = root2.Left, root2.Right
	}

	node.Val = lVal + rVal
	node.Left = mergeTrees(ll, rl)
	node.Right = mergeTrees(lr, rr)

	return node
}
