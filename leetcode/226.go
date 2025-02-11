package leetcode

func invertTree(root *TreeNode) *TreeNode {
	var invert func(root *TreeNode)
	invert = func(root *TreeNode) {
		if root == nil {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
	}
	invert(root)
	return root
}
