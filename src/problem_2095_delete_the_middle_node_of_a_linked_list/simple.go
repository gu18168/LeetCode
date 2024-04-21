package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	fakeHead := &ListNode{Next: head}
	oneStep, twoStep := fakeHead, fakeHead

	for twoStep.Next != nil && twoStep.Next.Next != nil {
		oneStep = oneStep.Next
		twoStep = twoStep.Next.Next
	}

	oneStep.Next = oneStep.Next.Next

	return fakeHead.Next
}
