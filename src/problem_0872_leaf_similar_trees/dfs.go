package solution

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	result1 := make([]int, 0)
	result2 := make([]int, 0)

	result1 = dfs(root1, result1)
	result2 = dfs(root2, result2)

	if len(result1) != len(result2) {
		return false
	}

	for i, val := range result1 {
		if result2[i] != val {
			return false
		}
	}

	return true
}

func dfs(node *TreeNode, result []int) []int {
	if node.Left == nil && node.Right == nil {
		return append(result, node.Val)
	}

	if node.Left != nil {
		result = dfs(node.Left, result)
	}

	if node.Right != nil {
		result = dfs(node.Right, result)
	}

	return result
}
