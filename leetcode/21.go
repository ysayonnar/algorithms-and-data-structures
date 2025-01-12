package leetcode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return list1
	} else if list1 == nil && list2 != nil {
		return list2
	} else if list1 != nil && list2 == nil {
		return list1
	}

	head := &ListNode{}
	headPtr := head

	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			headPtr.Next = list2
			headPtr = headPtr.Next
			list2 = list2.Next
		} else {
			headPtr.Next = list1
			headPtr = headPtr.Next
			list1 = list1.Next
		}
	}
	if list1 == nil {
		headPtr.Next = list2
	} else {
		headPtr.Next = list1
	}

	return head.Next
}
