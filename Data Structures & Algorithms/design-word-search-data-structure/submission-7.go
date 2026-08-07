type WordDictionary struct {
    Children map[byte]*WordDictionary
	IsEnd bool
}

func Constructor() WordDictionary {
    return WordDictionary{
		Children: map[byte]*WordDictionary{},
		IsEnd: false,
	}
}

func (this *WordDictionary) AddWord(word string)  {
	clone := this

	for i := 0; i < len(word); i++ {
		cur := word[i]

		if _, found := clone.Children[cur]; !found {
			clone.Children[cur] = &WordDictionary{
				Children: map[byte]*WordDictionary{},
				IsEnd: false,
			}
		}

		clone = clone.Children[cur]
	}

	clone.IsEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	cur := this 

	for i := 0; i < len(word); i++ {
		w := word[i]

		if w == '.' {
			for _, c := range cur.Children {
				if c.Search(word[i+1:]) {
					return true
				}
			}

			return false
		}
		
		if _, found := cur.Children[w]; !found {
			return false
		}

		cur = cur.Children[w]
	}

	return cur.IsEnd
}
