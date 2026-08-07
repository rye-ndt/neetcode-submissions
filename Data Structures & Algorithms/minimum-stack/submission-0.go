// use 2 stacks 
// one to store the real stack 
// one to store the minimum snapshot at each step 

type MinStack struct {
	Stack []int
	Min []int	
}

func Constructor() MinStack {
	return MinStack{
		Stack: []int{},
		Min: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)

	if len(this.Min) == 0 {
		this.Min = append(this.Min, val)
		return
	}

	curMin := this.Min[len(this.Min) - 1]

	if val < curMin {
		this.Min = append(this.Min, val)
		return 
	} 

	this.Min = append(this.Min, curMin)
}

func (this *MinStack) Pop() {
	if len(this.Stack) == 0 {
		return 
	}

	this.Stack = this.Stack[:len(this.Stack) - 1]
	this.Min = this.Min[:len(this.Min) - 1]
}

func (this *MinStack) Top() int {
	if len(this.Stack) == 0 {
		return 0
	}

	return this.Stack[len(this.Stack) - 1]
}

func (this *MinStack) GetMin() int {
	if len(this.Min) == 0 {
		return 0
	}
	
	return this.Min[len(this.Min) - 1]
}	
