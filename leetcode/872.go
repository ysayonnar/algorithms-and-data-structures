package leetcode

func preorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	} else if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	} else if root.Left == nil && root.Right != nil {
		return append([]int{}, preorder(root.Right)...)
	} else if root.Left != nil && root.Right == nil {
		return append([]int{}, preorder(root.Left)...)
	} else {
		return append(preorder(root.Left), preorder(root.Right)...)
	}
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 := preorder(root1)
	leaf2 := preorder(root2)
	if len(leaf1) != len(leaf2) {
		return false
	}

	for i := 0; i < len(leaf1); i++ {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}
	return true
}
