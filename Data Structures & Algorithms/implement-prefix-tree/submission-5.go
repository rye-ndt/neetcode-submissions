type Trie struct {
	children map[rune]*Trie
	isEnd bool
}

func Constructor() *Trie {
	return &Trie{
		children: map[rune]*Trie{},
		isEnd: false,
	}
}

func (this *Trie) Insert(word string) {
	clone := this 

	for _, w := range word {
		if clone.children[w] == nil {
			clone.children[w] = Constructor()
		}

		clone = clone.children[w]
	}

	clone.isEnd = true
}

func (this *Trie) Search(word string) bool {
	clone := this 

	for _, w := range word {
		if clone.children[w] == nil { return false }
		clone = clone.children[w]
	}

	return clone.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	clone := this 

	for _, w := range prefix {
		if clone.children[w] == nil { return false }
		clone = clone.children[w]
	}

	return true
}
