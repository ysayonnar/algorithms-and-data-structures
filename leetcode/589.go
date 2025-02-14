package leetcode

type Node3 struct {
	Val      int
	Children []*Node3
}

func preorder3(root *Node3) []int {
	result := []int{}

	var dfs func(root *Node3)
	dfs = func(root *Node3) {
		if root == nil {
			return
		}

		result = append(result, root.Val)

		for _, children := range root.Children {
			dfs(children)
		}
	}
	dfs(root)
	return result
}
