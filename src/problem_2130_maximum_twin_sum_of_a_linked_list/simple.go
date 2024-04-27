package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	halfN := 1
	oneStep, twoStep := head, head.Next

	for twoStep.Next != nil {
		oneStep = oneStep.Next
		twoStep = twoStep.Next.Next
		halfN += 1
	}

	var halfHead *ListNode

	for oneStep.Next != nil {
		newHead := &ListNode{Val: oneStep.Next.Val}
		newHead.Next = halfHead

		halfHead = newHead
		oneStep = oneStep.Next
	}

	result := 0
	frontCurrent, backCurrent := head, halfHead

	for i := 0; i < halfN; i++ {
		result = max(result, frontCurrent.Val+backCurrent.Val)

		frontCurrent = frontCurrent.Next
		backCurrent = backCurrent.Next
	}

	return result
}
