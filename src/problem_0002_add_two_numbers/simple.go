package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	now := &head

	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		val := carry
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		carry, val = val/10, val%10

		*now = &ListNode{Val: val}
		now = &(*now).Next
	}

	return head
}
