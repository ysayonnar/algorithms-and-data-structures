package leetcode

func dfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	result := []int{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current.Val)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}

	return result
}

func insertOne(root *TreeNode, value int) {
	node := &TreeNode{Val: value}
	if root == nil {
		root = node
		return
	}

	current := root
	for {
		if value < current.Val {
			if current.Left == nil {
				current.Left = node
				return
			}
			current = current.Left
		} else if value > current.Val {
			if current.Right == nil {
				current.Right = node
				return
			}
			current = current.Right
		} else {
			return
		}
	}
}

func insertAll(values []int) *TreeNode {
	root := &TreeNode{Val: values[0]}
	for _, value := range values[1:] {
		insertOne(root, value)
	}

	return root
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val == key && root.Left == nil && root.Right == nil {
		return nil
	}

	current := root
	var prev *TreeNode
	for current != nil {
		if key == current.Val {
			valuesLeft := dfs(current.Left)
			valuesRight := dfs(current.Right)
			values := append(valuesLeft, valuesRight...)
			if len(values) != 0 {
				node := insertAll(values)
				if prev == nil {
					return node
				}
				if key < prev.Val {
					prev.Left = node
				} else {
					prev.Right = node
				}
			} else {
				if key < prev.Val {
					prev.Left = nil
				} else {
					prev.Right = nil
				}
			}
			break
		} else if key < current.Val {
			prev = current
			current = current.Left
		} else {
			prev = current
			current = current.Right
		}
	}

	return root
}
