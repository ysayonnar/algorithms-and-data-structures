package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	current := head
	middle := head
	i := 0
	j := 0

	for current.Next != nil {
		if i/2 != j {
			j++
			middle = middle.Next
		}
		i++
		current = current.Next
	}
	middle.Next = middle.Next.Next
	return head
}
