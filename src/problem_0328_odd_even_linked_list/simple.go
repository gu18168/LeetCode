package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	oddStep := head

	if head.Next == nil {
		return head
	}

	evenHead, evenStep := head.Next, head.Next

	for evenStep != nil && evenStep.Next != nil {
		oddStep.Next = evenStep.Next
		oddStep = oddStep.Next

		evenStep.Next = oddStep.Next
		evenStep = evenStep.Next
	}

	oddStep.Next = evenHead

	return head
}
