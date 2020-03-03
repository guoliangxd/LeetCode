package leetcode

/*Implement a MyQueue class which implements a queue using two stacks.*/

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.length == 0 {
		return -1
	}
	for this.S1.length != 1 {
		this.S2.Push(this.S1.Pop())
	}
	rlt := this.S1.Pop()
	this.S1.Push(rlt)
	for this.S2.length != 0 {
		this.S1.Push(this.S2.Pop())
	}
	return rlt
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.length == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
type Stack struct {
	data   []int
	length int
}

func NewStack() *Stack {
	return &Stack{
		make([]int, 0, 1),
		0,
	}
}

func (this *Stack) IsEmpty() bool {
	return this.length == 0
}

func (this *Stack) Push(val int) {
	this.data = append(this.data, val)
	this.length++
}

func (this *Stack) Pop() int {
	if this.IsEmpty() {
		return -1
	} else {
		rlt := this.data[this.length-1]
		this.data = this.data[:this.length-1]
		this.length--
		return rlt
	}
}

type MyQueue struct {
	S1     *Stack
	S2     *Stack
	length int
}

func Constructor() MyQueue {
	return MyQueue{
		NewStack(),
		NewStack(),
		0,
	}
}

func (this *MyQueue) Push(x int) {
	this.S1.Push(x)
	this.length++
}

func (this *MyQueue) Pop() int {
	if this.length == 0 {
		return -1
	}
	for this.S1.length != 1 {
		this.S2.Push(this.S1.Pop())
	}
	rlt := this.S1.Pop()
	for this.S2.length != 0 {
		this.S1.Push(this.S2.Pop())
	}
	this.length--
	return rlt
}
