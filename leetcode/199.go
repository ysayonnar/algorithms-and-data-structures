package leetcode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			current := queue[0]
			queue = queue[1:]
			if i == l-1 {
				result = append(result, current.Val)
			}
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
	}

	return result
}
