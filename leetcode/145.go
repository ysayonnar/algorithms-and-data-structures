package leetcode

func postorderTraversal(root *TreeNode) []int {
	result := []int{}

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		dfs(root.Right)
		result = append(result, root.Val)
	}
	dfs(root)

	return result
}
