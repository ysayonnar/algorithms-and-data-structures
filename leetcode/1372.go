package leetcode

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxZigZag(root *TreeNode, direction bool) int {
	if root == nil {
		return -1
	}
	if !direction {
		return 1 + maxZigZag(root.Left, true)
	} else {
		return 1 + maxZigZag(root.Right, false)
	}
}

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := maxZigZag(root, false)
	r := maxZigZag(root, true)

	longestSubTree := max(longestZigZag(root.Left), longestZigZag(root.Right))

	return max(max(l, r), longestSubTree)
}
