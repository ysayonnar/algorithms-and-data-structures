package leetcode2

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		n := len(queue)
		level := []int{}
		for i := 0; i < n; i++ {
			current := queue[0]
			queue = queue[1:]
			level = append(level, current.Val)

			for _, node := range current.Children {
				queue = append(queue, node)
			}
		}

		result = append(result, level)
	}

	return result
}
