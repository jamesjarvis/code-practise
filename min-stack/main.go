import "math"

type MinStackItem struct {
    val int
    min int
}

type MinStack struct {
    vec []MinStackItem
}

/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(x int)  {
    currentMin := math.MaxInt32
    if(len(this.vec) != 0) {
        currentMin = this.vec[len(this.vec)-1].min
    }
    temp := MinStackItem{
        val: x,
        min: int(math.Min(float64(currentMin),float64(x))),
    }
    this.vec = append(this.vec, temp)
}


func (this *MinStack) Pop()  {
    this.vec = this.vec[:len(this.vec)-1]
}


func (this *MinStack) Top() int {
    return this.vec[len(this.vec)-1].val
}


func (this *MinStack) GetMin() int {
    return this.vec[len(this.vec)-1].min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */