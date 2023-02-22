package solution

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	nodes := []*Node{root}
	nextNodes := make([]*Node, 0)

	for len(nodes) != 0 {
		node := nodes[0]
		nodes = nodes[1:]

		if node.Left != nil {
			nextNodes = append(nextNodes, node.Left)
		}

		if node.Right != nil {
			nextNodes = append(nextNodes, node.Right)
		}

		if len(nodes) == 0 {
			nodes = nextNodes
			nextNodes = make([]*Node, 0)

			continue
		}

		node.Next = nodes[0]
	}

	return root
}
