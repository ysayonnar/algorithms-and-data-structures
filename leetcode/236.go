package leetcode

func isNodeInTree(root, p *TreeNode) bool {
	if root == nil {
		return false
	}
	if root == p {
		return true
	}

	return isNodeInTree(root.Left, p) || isNodeInTree(root.Right, p)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if !isNodeInTree(root, p) || !isNodeInTree(root, q) {
		return nil
	}

	lca1 := lowestCommonAncestor(root.Left, p, q)
	lca2 := lowestCommonAncestor(root.Right, p, q)
	if lca1 == nil && lca2 == nil {
		return root
	} else if lca1 != nil {
		return lca1
	} else {
		return lca2
	}
}
