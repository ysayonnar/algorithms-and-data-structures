package leetcode

func countNodesFromLeft(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ctr := 1
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		} else {
			ctr++
		}

		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root.Left)
	return ctr
}

func kthSmallest(root *TreeNode, k int) int {
	var nodes int
	current := root
	for nodes != k {
		nodes = countNodesFromLeft(current)
		if k < nodes {
			current = current.Left
		} else if k > nodes {
			current = current.Right
		}
	}

	return current.Val
}
