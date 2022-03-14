type MyQueue struct {
    items []int
}


func Constructor() MyQueue {
    return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
    this.items = append(this.items,x)
}


func (this *MyQueue) Pop() int {
    length:= len(this.items)
    if length==1{
        toRemove := this.items[0]
        this.items = this.items[:0]
        return toRemove
    }else{
        toRemove := this.items[0]
        this.items = this.items[1:length]
        return toRemove
    }
}


func (this *MyQueue) Peek() int {
    return this.items[0]
}


func (this *MyQueue) Empty() bool {
    if len(this.items)==0{
        return true
    }   else{
        return false   
    }
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
