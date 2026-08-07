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
	clone := this // preserve 

	for i := 0; i < len(word); i++ {
		cur := word[i]

		if _, found := clone.Children[cur]; !found {
			clone.Children[cur] = &PrefixTree{
				Children: map[byte]*PrefixTree{},
				IsEnd: false,
			}
		}

		clone = clone.Children[cur]
	}

	clone.IsEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		cur := word[i]

		if _, found := this.Children[cur]; !found {
			return false
		}

		this = this.Children[cur]
	}

	return this.IsEnd
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		cur := prefix[i]

		if _, found := this.Children[cur]; !found {
			return false
		}

		this = this.Children[cur]
	}

	return true
}
