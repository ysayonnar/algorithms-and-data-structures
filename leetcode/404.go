package leetcode

func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0

	var count func(root *TreeNode, isLeft bool)
	count = func(root *TreeNode, isLeft bool) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil && isLeft {
			sum += root.Val
			return
		}

		count(root.Left, true)
		count(root.Right, false)
	}

	count(root, false)

	return sum
}
