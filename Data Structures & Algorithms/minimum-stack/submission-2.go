type MinStack struct {
	vals []int
	mins []int
}

func Constructor() MinStack {
	return MinStack{
		vals: []int{},
		mins: []int{},
	}
}

func (t *MinStack) Push(val int) {
	t.vals = append(t.vals, val)

	if len(t.mins) == 0 {
		t.mins = append(t.mins, val)
	} else {
		t.mins = append(t.mins, min(t.mins[len(t.mins)-1], val))
	}
}

func (t *MinStack) Pop() {
	if len(t.vals) == 0 { return }

	t.vals = t.vals[:len(t.vals)-1]
	t.mins = t.mins[:len(t.mins)-1]
}

func (t *MinStack) Top() int {
	return t.vals[len(t.vals)-1]
}

func (t *MinStack) GetMin() int {
	return t.mins[len(t.mins)-1]
}	
