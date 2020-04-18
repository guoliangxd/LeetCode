package leetcode

/*在战略游戏中，玩家往往需要发展自己的势力来触发各种新的剧情。一个势力的主要属性有三种，分别是文明等级（C），资源储备（R）以及人口数量（H）。在游戏开始时（第 0 天），三种属性的值均为 0。

随着游戏进程的进行，每一天玩家的三种属性都会对应增加，我们用一个二维数组 increase 来表示每天的增加情况。这个二维数组的每个元素是一个长度为 3 的一维数组，例如 [[1,2,1],[3,4,2]] 表示第一天三种属性分别增加 1,2,1 而第二天分别增加 3,4,2。

所有剧情的触发条件也用一个二维数组 requirements 表示。这个二维数组的每个元素是一个长度为 3 的一维数组，对于某个剧情的触发条件 c[i], r[i], h[i]，如果当前 C >= c[i] 且 R >= r[i] 且 H >= h[i] ，则剧情会被触发。

根据所给信息，请计算每个剧情的触发时间，并以一个数组返回。如果某个剧情不会被触发，则该剧情对应的触发时间为 -1 。*/

import "sort"

type Array [][]int

func (this Array) Len() int {
	return len(this)
}

func (this Array) Less(i, j int) bool {
	if this[i][0] < this[j][0] {
		return true
	} else if this[i][0] > this[j][0] {
		return false
	} else {
		if this[i][1] < this[j][1] {
			return true
		} else if this[i][1] > this[i][1] {
			return false
		} else {
			return this[i][2] < this[i][2]
		}
	}
}

func (this Array) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func getTriggerTime(increase [][]int, requirements [][]int) []int {
	today := [3]int{}
	days := len(increase)
	size := len(requirements)
	res := make([]int, size)
	sort.Sort(Array(requirements))

	for i := 0; i < size; i++ {
		res[i] = -1
	}

	for i := 0; i < days; i++ {
		today[0] += increase[i][0]
		today[1] += increase[i][1]
		today[2] += increase[i][2]

		for j := 0; j < size; j++ {
			if res[j] != -1 {
				continue
			}

			if requirements[j][0] > today[0] {
				break
			}

			if requirements[j][1] <= today[1] && requirements[j][2] <= today[2] {
				res[j] = i + 1
			}
		}
	}
	return res
}
