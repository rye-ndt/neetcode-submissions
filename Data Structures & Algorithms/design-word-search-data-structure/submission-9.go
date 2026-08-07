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
	for _, w := range word {
		cur := byte(w)

		if _, found := this.Children[cur]; !found {
			this.Children[cur] = &WordDictionary{
				Children: map[byte]*WordDictionary{},
				IsEnd: false,
			}
		}	 

		this = this.Children[cur]
	}

	this.IsEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		cur := word[i]

		if cur == '.' {
			// recursive
			for _, c := range this.Children {
				if c.Search(word[i+1:]) {
					return true
				}
			}

			return false
		}

		// if not, find as normal 
		if _, found := this.Children[cur]; !found {
			return false
		}

		this = this.Children[cur]
	}

	return this.IsEnd
}
