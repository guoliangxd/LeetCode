package leetcode

/*用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，
分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

type Stack struct {
    data []int
    length int
}

func NewStack() *Stack {
    return &Stack{
        make([]int, 0, 1),
        0,
    }
}

func (this *Stack)IsEmpty() bool {
    return this.length == 0
}

func (this *Stack)Push(val int) {
    this.data = append(this.data, val)
    this.length++
}

func (this *Stack)Pop() int {
    if this.IsEmpty() {
        return -1
    } else {
        rlt := this.data[this.length - 1]
        this.data = this.data[ : this.length - 1]
        this.length--
        return rlt
    }
}

type CQueue struct {
    S1 *Stack
    S2 *Stack
    length int
}


func Constructor() CQueue {
    return CQueue {
        NewStack(),
        NewStack(),
        0,
    }
}


func (this *CQueue) AppendTail(value int)  {
    this.S1.Push(value)
    this.length++
}


func (this *CQueue) DeleteHead() int {
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


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */