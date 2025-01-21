package leetcode

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	node := &TreeNode{Val: val}
	if root == nil {
		return node
	}

	current := root
	for current != nil {
		if val < current.Val {
			if current.Left == nil {
				current.Left = node
				break
			} else {
				current = current.Left
			}
		} else if val > current.Val {
			if current.Right == nil {
				current.Right = node
				break
			} else {
				current = current.Right
			}
		}

	}

	return root
}
