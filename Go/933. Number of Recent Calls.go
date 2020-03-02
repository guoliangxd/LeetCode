package leetcode

/*Write a class RecentCounter to count recent requests.
It has only one method: ping(int t), where t represents some time in milliseconds.
Return the number of pings that have been made from 3000 milliseconds ago until now.
Any ping with time in [t - 3000, t] will count, including the current ping.
It is guaranteed that every call to ping uses a strictly larger value of t than before.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-recent-calls
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

type RecentCounter struct {
	data []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		make([]int, 0, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.data = append(this.data, t)
	length := len(this.data)
	now := this.data[length-1]
	index := 0
	for i := 0; i < length; i++ {
		if now-this.data[i] <= 3000 {
			index = i
			break
		}
	}
	if index > 0 {
		this.data = this.data[index-1:]
	}
	return length - index
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
