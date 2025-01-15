package leetcode

func bypass(root *TreeNode, path []int) int {
	acc := 0
	if root == nil {
		return 0
	} else if len(path) == 0 {
		path = append(path, root.Val)
		acc++
	} else if path[len(path)-1] <= root.Val {
		path = append(path, root.Val)
		acc++
	}
	acc += bypass(root.Left, path)
	acc += bypass(root.Right, path)
	return acc

}

func goodNodes(root *TreeNode) int {
	path := []int{}
	return bypass(root, path)
}
