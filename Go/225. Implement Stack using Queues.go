package leetcode

/*Implement the following operations of a stack using queues.
push(x) -- Push element x onto stack.
pop() -- Removes the element on top of the stack.
top() -- Get the top element.
empty() -- Return whether the stack is empty.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-stack-using-queues
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/
type MyStack struct {
	Q1     *MyQueue
	Q2     *MyQueue
	Length int
}

type MyQueue struct {
	data   []int
	Length int
}

//获取队首元素
func (this *MyQueue) Get() int {
	if this.Length > 0 {
		this.Length--
		rlt := this.data[0]
		this.data = this.data[1:]
		return rlt
	} else {
		return -1
	}
}

//在队尾添加元素
func (this *MyQueue) Put(value int) {
	this.data = append(this.data, value)
	this.Length++
}

//队列是否为空
func (this *MyQueue) IsEmpty() bool {
	if this.Length == 0 {
		return true
	} else {
		return false
	}
}

func NewQueue() *MyQueue {
	return &MyQueue{
		make([]int, 0, 1),
		0,
	}
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		NewQueue(),
		NewQueue(),
		0,
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.Q1.Put(x)
	this.Length++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for this.Q1.Length != 1 {
		this.Q2.Put(this.Q1.Get())
	}
	rlt := this.Q1.Get()
	this.Q1, this.Q2 = this.Q2, this.Q1
	return rlt
}

/** Get the top element. */
func (this *MyStack) Top() int {
	rlt := this.Pop()
	this.Push(rlt)
	return rlt
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Length == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
