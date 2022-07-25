package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}

	for root != p && root != q {
		if root.Val > p.Val && root.Val < q.Val {
			break
		} else if root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val {
			root = root.Right
		}
	}

	return root
}
