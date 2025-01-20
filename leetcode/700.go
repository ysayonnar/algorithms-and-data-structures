package leetcode

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	current := root
	for current != nil {
		if val == current.Val {
			return current
		} else if val < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return nil
}
