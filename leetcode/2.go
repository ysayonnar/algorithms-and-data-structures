package leetcode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	perenos := 0
	result := &ListNode{}
	resultPtr := result

	for l1 != nil || l2 != nil {
		if l1 == nil {
			for l2 != nil {
				resultPtr.Val = l2.Val + perenos
				if resultPtr.Val > 9 {
					resultPtr.Val -= 10
					perenos = 1
				} else {
					perenos = 0
				}
				if l2.Next != nil {
					resultPtr.Next = &ListNode{}
					resultPtr = resultPtr.Next
				}
				l2 = l2.Next
			}
			break
		} else if l2 == nil {
			for l1 != nil {
				resultPtr.Val = l1.Val + perenos
				if resultPtr.Val > 9 {
					resultPtr.Val -= 10
					perenos = 1
				} else {
					perenos = 0
				}
				if l1.Next != nil {
					resultPtr.Next = &ListNode{}
					resultPtr = resultPtr.Next
				}
				l1 = l1.Next
			}
			break
		} else {
			sum := l1.Val + l2.Val + perenos
			if sum > 9 {
				sum -= 10
				perenos = 1
			} else {
				perenos = 0
			}
			resultPtr.Val = sum
			l1 = l1.Next
			l2 = l2.Next
			if l1 != nil || l2 != nil {
				resultPtr.Next = &ListNode{}
				resultPtr = resultPtr.Next
			}
		}
	}
	if perenos == 1 {
		resultPtr.Next = &ListNode{Val: perenos}
	}

	return result
}
