package leetcode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	var prev *TreeNode

	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Right)
		dfs(node.Left)

		node.Right = prev
		node.Left = nil
		prev = node
	}

	dfs(root)
}
