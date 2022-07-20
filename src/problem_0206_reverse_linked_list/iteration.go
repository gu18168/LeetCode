package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	now := head

	for now != nil {
		next := now.Next
		now.Next = prev
		prev, now = now, next
	}

	return prev
}