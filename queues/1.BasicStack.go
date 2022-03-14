type MyStack struct {
	items []int
}

func Constructor() MyStack {
	return MyStack{}
}
func (this *MyStack) Push(x int) {
	this.items = append(this.items, x)
}
func (this *MyStack) Pop() int {
	length := len(this.items)
	if length != 1 {
		removed := this.items[length-1]
		this.items = this.items[:length-1]
		return removed
	} else {
		removed := this.items[0]
		this.items = this.items[:0]
		return removed
	}

}

func (this *MyStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MyStack) Empty() bool {
	length := len(this.items)
	if length != 0 {
		return false
	} else {
		return true
	}

}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */