package solution

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}

	result := []int{root.Val}

	for _, child := range root.Children {
		subResult := preorder(child)
		result = append(result, subResult...)
	}

	return result
}
