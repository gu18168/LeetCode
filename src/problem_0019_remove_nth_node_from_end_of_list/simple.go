package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left, right := &head, head

	for i := 0; i < n; i++ {
		right = right.Next
	}

	for right != nil {
		right = right.Next
		left = &(*left).Next
	}

	*left = (*left).Next

	return head
}
