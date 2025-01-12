package dataStructures

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *Tree) Insert(value int) {
	node := &TreeNode{Value: value}
	if t.Root == nil {
		t.Root = node
		return
	}

	current := t.Root
	for {
		if value < current.Value {
			if current.Left == nil {
				current.Left = node
				return
			}
			current = current.Left
		} else if value > current.Value {
			if current.Right == nil {
				current.Right = node
				return
			}
			current = current.Right
		} else {
			return // value уже в дереве
		}
	}
}

func (t *Tree) Search(value int) bool {
	current := t.Root
	for current != nil {
		if current.Value == value {
			return true
		} else if value < current.Value {
			current = current.Left
		} else if value > current.Value {
			current = current.Right
		}
	}
	return false
}

func (t *Tree) Delete(value int) {
	current := t.Root
	prev := t.Root

	for current != nil {
		if value == current.Value {
			if value < prev.Value {
				prev.Left = nil
			} else {
				prev.Right = nil
			}
			return
		} else if value < current.Value {
			prev = current
			current = current.Left
		} else if value > current.Value {
			prev = current
			current = current.Right
		}
	}
}

func (t *Tree) PrintTreeStructure() {
	printTreeStructure(t.Root, 0, "")
}

func printTreeStructure(node *TreeNode, level int, prefix string) {
	if node == nil {
		return
	}

	indent := strings.Repeat("  ", level)
	fmt.Printf("%s%s%d\n", indent, prefix, node.Value)

	printTreeStructure(node.Left, level+1, "L: ")
	printTreeStructure(node.Right, level+1, "R: ")
}
