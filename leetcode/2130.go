package leetcode

func reverse(head *ListNode) *ListNode {
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

func pairSum(head *ListNode) int {
	middle := head
	fast := head

	for fast != nil {
		fast = fast.Next.Next
		middle = middle.Next
	}
	middle = reverse(middle)

	i, j := head, middle
	sum := i.Val + j.Val
	for i.Next != nil && j.Next != nil {
		i = i.Next
		j = j.Next
		if i.Val+j.Val > sum {
			sum = i.Val + j.Val
		}
	}

	return sum
}
