package leetcode

/*We are given a matrix with R rows and C columns has cells with integer coordinates (r, c), where 0 <= r < R and 0 <= c < C.

Additionally, we are given a cell in that matrix with coordinates (r0, c0).

Return the coordinates of all cells in the matrix, sorted by their distance from (r0, c0) from smallest distance to largest distance.  Here, the distance between two cells (r1, c1) and (r2, c2) is the Manhattan distance, |r1 - r2| + |c1 - c2|.  (You may return the answer in any order that satisfies this condition.)

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/matrix-cells-in-distance-order
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import (
	"math"
	"sort"
)

type node struct {
	loc []int
	dis int
}

type NodeList []node

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	nodeList := make(NodeList, 0, R*C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			nodeList = append(nodeList, node{[]int{i, j}, getManhattan(r0, c0, i, j)})
		}
	}
	sort.Sort(nodeList)
	result := make([][]int, 0)
	for i := 0; i < nodeList.Len(); i++ {
		result = append(result, nodeList[i].loc)
	}
	return result
}

func getManhattan(r0, c0, r, c int) int {
	return int(math.Abs(float64(r0-r))) + int(math.Abs(float64(c0-c)))
}

func (this NodeList) Len() int {
	return len(this)
}

func (this NodeList) Less(i, j int) bool {
	return this[i].dis < this[j].dis
}

func (this NodeList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
