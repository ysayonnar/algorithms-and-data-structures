package leetcode2

type ListNode struct {
	Val  int
	Next *ListNode
}

// bad memory complexity
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	nodes := map[*ListNode]struct{}{}

	current := headA
	for current != nil {
		nodes[current] = struct{}{}
		current = current.Next
	}

	current = headB
	for current != nil {
		if _, ok := nodes[current]; ok {
			return current
		}
		current = current.Next
	}

	return nil
}

// bad memory complexity
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	nodes := map[*ListNode]struct{}{}

	current := headA
	for current != nil {
		nodes[current] = struct{}{}
		current = current.Next
	}

	current = headB
	for current != nil {
		if _, ok := nodes[current]; ok {
			return current
		}
		current = current.Next
	}

	return nil
}

// best solution: O(m+n) - time, O(1) - memory
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := 0
	lenB := 0

	current := headA
	for current != nil {
		lenA++
		current = current.Next
	}

	current = headB
	for current != nil {
		lenB++
		current = current.Next
	}

	skip := 0
	if lenA > lenB {
		skip = lenA - lenB

		current = headA
		for range skip {
			current = current.Next
		}
		headA = current
	} else if lenA < lenB {
		skip = lenB - lenA

		current = headB
		for range skip {
			current = current.Next
		}
		headB = current
	}

	currentA := headA
	currentB := headB
	for currentA != nil && currentB != nil {
		if currentA == currentB {
			return currentA
		}
		currentA = currentA.Next
		currentB = currentB.Next
	}

	return nil
}
