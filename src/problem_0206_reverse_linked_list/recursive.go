package solution

func reverseListRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	revHead := reverseListRec(head.Next)
	head.Next.Next = head
	head.Next = nil

	return revHead
}

