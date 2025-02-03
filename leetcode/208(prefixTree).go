package leetcode

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

func (this *Trie) Search(word string) bool {
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

func (this *Trie) StartsWith(prefix string) bool {
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
