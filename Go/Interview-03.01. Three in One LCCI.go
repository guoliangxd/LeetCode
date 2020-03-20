package leetcode

/*Describe how you could use a single array to implement three stacks.

Yout should implement push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、peek(stackNum) methods. stackNum is the index of the stack. value is the value that pushed to the stack.

The constructor requires a stackSize parameter, which represents the size of each stack.*/

type TripleInOne struct {
	data []int

	size int

	index1 int
	index2 int
	index3 int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		make([]int, stackSize*3),
		stackSize,
		0,
		stackSize,
		stackSize + stackSize,
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	switch stackNum {
	case 0:
		if this.index1 >= this.size {
			return
		} else {
			this.data[this.index1] = value
			this.index1++
		}
	case 1:
		if this.index2 >= 2*this.size {
			return
		} else {
			this.data[this.index2] = value
			this.index2++
		}
	case 2:
		if this.index3 >= 3*this.size {
			return
		} else {
			this.data[this.index3] = value
			this.index3++
		}
	default:
	}
}

func (this *TripleInOne) Pop(stackNum int) int {
	switch stackNum {
	case 0:
		if this.index1 == 0 {
			return -1
		} else {
			this.index1--
			return this.data[this.index1]
		}
	case 1:
		if this.index2 == this.size {
			return -1
		} else {
			this.index2--
			return this.data[this.index2]
		}
	case 2:
		if this.index3 == 2*this.size {
			return -1
		} else {
			this.index3--
			return this.data[this.index3]
		}
	default:
		return -1
	}
	return -1
}

func (this *TripleInOne) Peek(stackNum int) int {
	switch stackNum {
	case 0:
		if this.index1 == 0 {
			return -1
		} else {
			return this.data[this.index1-1]
		}
	case 1:
		if this.index2 == this.size {
			return -1
		} else {
			return this.data[this.index2-1]
		}
	case 2:
		if this.index3 == 2*this.size {
			return -1
		} else {
			return this.data[this.index3-1]
		}
	default:
		return -1
	}
	return -1
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	switch stackNum {
	case 0:
		if this.index1 == 0 {
			return true
		} else {
			return false
		}
	case 1:
		if this.index2 == this.size {
			return true
		} else {
			return false
		}
	case 2:
		if this.index3 == 2*this.size {
			return true
		} else {
			return false
		}
	default:
		return false
	}
	return false
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */
