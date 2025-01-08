package leetcode

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	result := &ListNode{}

	for head.Next != nil {
		result.Val = head.Val
		result = &ListNode{Next: result}

		head = head.Next
	}
	result.Val = head.Val

	return result
}
