package leetcode

/*请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的时间复杂度都是O(1)。

若队列为空，pop_front 和 max_value 需要返回 -1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/dui-lie-de-zui-da-zhi-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

type MaxQueue struct {
	data []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.data) == 0 {
		return -1
	} else {
		max := this.data[0]
		for i := 1; i < len(this.data); i++ {
			if this.data[i] > max {
				max = this.data[i]
			}
		}
		return max
	}
}

func (this *MaxQueue) Push_back(value int) {
	this.data = append(this.data, value)
}

func (this *MaxQueue) Pop_front() int {
	rlt := -1
	if len(this.data) != 0 {
		rlt = this.data[0]
		this.data = this.data[1:]
	}
	return rlt
}
