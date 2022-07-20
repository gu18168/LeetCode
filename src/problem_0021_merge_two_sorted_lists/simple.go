package solution

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil {
        return list2
    }
    if list2 == nil {
        return list1
    }

    var head **ListNode
    var result *ListNode
    if list1.Val > list2.Val {
        head, result = &list2.Next, list2
        list2 = list2.Next
    } else {
        head, result = &list1.Next, list1
        list1 = list1.Next
    }

	for list1 != nil && list2 != nil {
		val1, val2 := list1.Val, list2.Val

		if val1 > val2 {
			*head = list2
			list2 = list2.Next
		} else {
			*head = list1
			list1 = list1.Next
		}

		head = &(*head).Next
	}

	if list1 != nil {
		*head = list1
	} else if list2 != nil {
		*head = list2
	}

	return result
}