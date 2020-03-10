package leetcode

/*How would you design a stack which, in addition to push and pop, has a function min which returns the minimum element? Push, pop and min should all operate in 0(1) time.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/min-stack-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

type MinStack struct {
	data    []int
	minData []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.minData) == 0 || x <= this.minData[len(this.minData)-1] {
		this.minData = append(this.minData, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.data) == 0 {
		return
	}
	temp := this.data[len(this.data)-1]
	this.data = this.data[:len(this.data)-1]
	if temp == this.minData[len(this.minData)-1] {
		this.minData = this.minData[:len(this.minData)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minData) == 0 {
		return -1
	}
	return this.minData[len(this.minData)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
