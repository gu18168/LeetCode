package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	now := []*TreeNode{root}
	result := make([][]int, 0)

	for len(now) != 0 {
		next := make([]*TreeNode, 0)
		layer := make([]int, 0, len(now))

		for _, node := range now {
			layer = append(layer, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}

		now = next
		result = append(result, layer)
	}

	return result
}
