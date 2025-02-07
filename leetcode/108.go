package leetcode

func MakeTree(root *TreeNode, nums []int) {
	if len(nums) == 0 {
		return
	}

	middle := len(nums) / 2
	root.Val = nums[middle]
	if middle > 0 {
		root.Left = &TreeNode{}
		MakeTree(root.Left, nums[:middle])
	}

	if middle < len(nums)-1 {
		root.Right = &TreeNode{}
		MakeTree(root.Right, nums[middle+1:])
	}
}

func sortedArrayToBST(nums []int) *TreeNode {
	root := TreeNode{}
	MakeTree(&root, nums)
	return &root
}
