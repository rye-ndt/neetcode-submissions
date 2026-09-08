type Dict struct {
	child map[rune]*Dict
	end bool
}

func Constructor() *Dict {
	return &Dict{
		child: map[rune]*Dict{},
		end: false,
	}
}

var New = Constructor

func (this *Dict) AddWord(word string)  {
	clone := this

	for _, w := range word {
		if clone.child[w] == nil {
			clone.child[w] = New()
		}

		clone = clone.child[w]
	}

	clone.end = true
}

func (this *Dict) Search(word string) bool {
	clone := this 

	for i, c := range word {
		if c == '.' {
			for _, childDict := range clone.child {
 				if childDict.Search(word[i+1:]) { return true }
			}

			return false
		} else {
			if clone.child[c] == nil { return false }
			clone = clone.child[c]
		}
	}

	return clone.end
}
