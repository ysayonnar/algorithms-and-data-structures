package leetcode

type Node1 struct {
	Children map[byte]*Node1
	isEnd    bool
}

type Trie1 struct {
	Root *Node1
}

func Constructor3() Trie {
	return Trie{Root: &Node{Children: make(map[byte]*Node), isEnd: true}}
}

func (this *Trie) Insert2(word string) {
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

func (this *Trie) Search2(word string) bool {
	current := this.Root
	symbols := []byte(word)
	for _, s := range symbols {
		if _, ok := current.Children[s]; !ok {
			return false
		} else {
			current = current.Children[s]
		}
	}
	return current.isEnd
}

func (this *Trie) StartsWith2(prefix string) bool {
	symbols := []byte(prefix)
	current := this.Root
	for _, s := range symbols {
		if _, ok := current.Children[s]; !ok {
			return false
		} else {
			current = current.Children[s]
		}
	}

	return true
}
