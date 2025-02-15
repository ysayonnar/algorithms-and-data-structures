package leetcode2

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	ctr := 0
	queue := []*Node{root}
	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			current := queue[0]
			queue = queue[1:]
			for _, node := range current.Children {
				queue = append(queue, node)
			}
		}
		ctr++
	}
	return ctr
}
