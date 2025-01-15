package leetcode

func reduce(root *TreeNode, targerSum int, reducer []int) int {
	if root == nil {
		return 0
	}

	sum := 0
	ctr := 0
	reducer = append(reducer, root.Val)

	for i := len(reducer) - 1; i >= 0; i-- {
		sum += reducer[i]
		if sum == targerSum {
			ctr++
		}
	}

	return ctr + reduce(root.Left, targerSum, reducer) + reduce(root.Right, targerSum, reducer)
}

func pathSum(root *TreeNode, targetSum int) int {
	arr := []int{}
	return reduce(root, targetSum, arr)
}
