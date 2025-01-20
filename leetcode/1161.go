package leetcode

func maxLevelSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	sum := root.Val
	currentLevel := 0
	maxLevel := 1

	for len(queue) > 0 {
		currentLevel++
		l := len(queue)
		tempSum := 0
		for i := 0; i < l; i++ {
			current := queue[0]
			queue = queue[1:]
			tempSum += current.Val

			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}

		if tempSum > sum {
			sum = tempSum
			maxLevel = currentLevel
		}
	}

	return maxLevel
}
