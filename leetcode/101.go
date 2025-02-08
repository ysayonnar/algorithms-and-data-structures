package leetcode

func isSymmetric(root *TreeNode) bool {
	left := []int{}
	right := []int{}

	var leftDfs func(root *TreeNode)
	var rightDfs func(root *TreeNode)

	leftDfs = func(root *TreeNode) {
		if root == nil {
			left = append(left, -101)
			return
		}
		left = append(left, root.Val)
		leftDfs(root.Left)
		leftDfs(root.Right)
	}

	rightDfs = func(root *TreeNode) {
		if root == nil {
			right = append(right, -101)
			return
		}
		right = append(right, root.Val)
		rightDfs(root.Right)
		rightDfs(root.Left)
	}

	leftDfs(root.Left)
	rightDfs(root.Right)

	if len(left) != len(right) {
		return false
	}

	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}
