package solution

type ListNode struct {
	Val int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head

	toggle := true
	for fast.Next != nil {
		fast = fast.Next

		if toggle {
			slow = slow.Next
		}
		toggle = !toggle
	}

	return slow
}