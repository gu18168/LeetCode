package solution

func detectCycleMap(head *ListNode) *ListNode {
	hasVisited := make(map[*ListNode]struct{})

	for head != nil {
		if _, ok := hasVisited[head]; ok {
			return head
		}

		hasVisited[head] = struct{}{}
		head = head.Next
	}

	return head
}
