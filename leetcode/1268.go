package leetcode

import "sort"

type Node struct {
	Children map[byte]*Node
	isEnd    bool
}

type Trie struct {
	Root *Node
}

func Constructor() Trie {
	return Trie{Root: &Node{Children: make(map[byte]*Node), isEnd: true}}
}

func (this *Trie) Insert(word string) {
	current := this.Root
	symbols := []byte(word)
	lastIndex := len(word)
	for i, s := range symbols {
		if _, ok := current.Children[s]; !ok {
			lastIndex = i
			break
		} else {
			current = current.Children[s]
		}
	}

	for i := lastIndex; i < len(word); i++ {
		current.Children[symbols[i]] = &Node{Children: make(map[byte]*Node)}
		current = current.Children[symbols[i]]
	}
	current.isEnd = true
}

func (this *Trie) StartsWith(prefix string) []string {
	result := []string{}
	symbols := []byte(prefix)
	current := this.Root
	for _, s := range symbols {
		if _, ok := current.Children[s]; !ok {
			return result
		} else {
			current = current.Children[s]
		}
	}

	var collect func(node *Node, p string)
	collect = func(node *Node, p string) {
		if node.isEnd {
			result = append(result, p)
		}
		for r, child := range node.Children {
			collect(child, p+string(r))
		}
	}

	collect(current, "")

	for i := 0; i < len(result); i++ {
		result[i] = prefix + result[i]
	}

	sort.Strings(result)

	if len(result) <= 3 {
		return result
	}

	return result[:3]
}

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := Constructor()
	for _, product := range products {
		trie.Insert(product)
	}

	result := [][]string{}
	for i := 0; i < len(searchWord); i++ {
		result = append(result, trie.StartsWith(searchWord[:i+1]))
	}

	return result
}
