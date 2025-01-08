package leetcode

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		return head
	}

	odd := head
	even := head.Next
	temp := even

	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}

	odd.Next = temp

	return head
}
