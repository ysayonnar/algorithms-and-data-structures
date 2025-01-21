package leetcode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	levels := [][]int{}

	for len(queue) > 0 {
		l := len(queue)
		level := []int{}
		for i := 0; i < l; i++ {
			current := queue[0]
			queue = queue[1:]
			level = append(level, current.Val)

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		levels = append(levels, level)
	}

	return levels
}
