package solution

func connectDFS(root *Node) *Node {
	if root == nil {
		return root
	}

	if root.Left != nil {
		root.Left.Next = root.Right
	}

	if root.Right != nil && root.Next != nil {
		root.Right.Next = root.Next.Left
	}

	connectDFS(root.Left)
	connectDFS(root.Right)

	return root
}
