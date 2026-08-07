type PrefixTree struct {
	Children map[byte]*PrefixTree
	IsEnd bool
}

func Constructor() PrefixTree {
    return PrefixTree{
		Children: map[byte]*PrefixTree{},
		IsEnd: false,
	}
}

func (this *PrefixTree) Insert(word string) {
	node := this 

	for i := 0; i < len(word); i++ {
		cur := word[i]

		if _, found := node.Children[cur]; !found {
			node.Children[cur] = &PrefixTree{
				Children: map[byte]*PrefixTree{},
				IsEnd: false,
			}
		}

		node = node.Children[cur]
	}

	node.IsEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		cur := word[i]

		if _, found := this.Children[cur]; found {
			this = this.Children[cur]
			continue
		} 

		return false
	}

	return this.IsEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		cur := prefix[i]

		if _, found := this.Children[cur]; found {
			this = this.Children[cur]
			continue
		}

		return false
	}

	return true
}
