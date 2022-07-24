package solution

func preorderIter(root *Node) []int {
	if root == nil {
		return nil
	}

	stack := []*Node{root}
	result := make([]int, 0)

	for len(stack) != 0 {
		l := len(stack)
		node := stack[l-1]
		stack = stack[:l-1]

		result = append(result, node.Val)
		for i := len(node.Children) - 1; i >=0; i-- {
			stack = append(stack, node.Children[i])
		}
	}

	return result
}
