package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := 1
	current := head

	for current != nil {
		length++
		current = current.Next
	}

	if length == n {
		return head.Next
	}

	i := 1
	current = head
	for current.Next != nil {
		if i == length-n {
			current.Next = current.Next.Next
			break
		}
		i++
		current = current.Next
	}

	return head
}
