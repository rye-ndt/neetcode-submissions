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
		if clone.child[w] == nil { clone.child[w] = New() }
		clone = clone.child[w]
	}

	clone.end = true
}

func (this *Dict) Search(word string) bool {
	clone := this 

	for i, c := range word {
		switch c {
			case '.':
				for _, child := range clone.child {
					if child.Search(word[i+1:]) { return true }
				}

				return false
			default: 
				if clone.child[c] == nil { return false }
				clone = clone.child[c]
		}
	}

	return clone.end
}
