package leetcode

import (
	"fmt"
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	result := []string{}

	var dfs func(root *TreeNode, path string)
	dfs = func(root *TreeNode, path string) {
		if root == nil {
			return
		}

		path += fmt.Sprintf("->%d", root.Val)
		if root.Left == nil && root.Right == nil {
			result = append(result, path)
		}
		dfs(root.Left, path)
		dfs(root.Right, path)
	}

	dfs(root.Left, strconv.Itoa(root.Val))
	dfs(root.Right, strconv.Itoa(root.Val))

	return result
}
