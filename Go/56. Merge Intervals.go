package leetcode

/*Given a collection of intervals, merge all overlapping intervals.*/

import "sort"

type Data [][]int

func (this Data) Less(i, j int) bool {
	return this[i][0] < this[j][0]
}

func (this Data) Len() int {
	return len(this)
}

func (this Data) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Sort(Data(intervals))
	for i := 0; i < len(intervals); i++ {
		if len(res) == 0 {
			res = append(res, intervals[i])
		} else {
			size := len(res)
			if res[size-1][1] >= intervals[i][0] {
				if res[size-1][1] < intervals[i][1] {
					res[size-1][1] = intervals[i][1]
				}
			} else {
				res = append(res, intervals[i])
			}
		}
	}
	return res
}
