package solution

type ListNode struct {
	Val int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
    slow, fast := head, head

    for fast != nil {
        slow = slow.Next
        if fast.Next == nil {
            return nil
        }
        fast = fast.Next.Next

		// 2(a + b) = a + n(b + c) + b -> a = c + (n - 1)(b + c)
		// so new `p` move a step, and `slow` move c step
        if fast == slow {
            p := head
            for p != slow {
                p = p.Next
                slow = slow.Next
            }
            return p
        }
    }

	return nil
}