package leetcode

/*小朋友 A 在和 ta 的小伙伴们玩传信息游戏，游戏规则如下：

有 n 名玩家，所有玩家编号分别为 0 ～ n-1，其中小朋友 A 的编号为 0
每个玩家都有固定的若干个可传信息的其他玩家（也可能没有）。传信息的关系是单向的（比如 A 可以向 B 传信息，但 B 不能向 A 传信息）。
每轮信息必须需要传递给另一个人，且信息可重复经过同一个人
给定总玩家数 n，以及按 [玩家编号,对应可传递玩家编号] 关系组成的二维数组 relation。返回信息从小 A (编号 0 ) 经过 k 轮传递到编号为 n-1 的小伙伴处的方案数；若不能到达，返回 0。*/

var relat map[int][]int
var cnt, max int

func numWays(n int, relation [][]int, k int) int {
	relat = make(map[int][]int)
	cnt = 0
	max = k
	for i := 0; i < len(relation); i++ {
		relat[relation[i][0]] = append(relat[relation[i][0]], relation[i][1])
	}

	search(0, n-1, 0)
	return cnt
}

func search(begin, end, step int) {
	if step == max {
		return
	}
	adj := relat[begin]
	step++
	for i := 0; i < len(adj); i++ {
		if adj[i] == end && step == max {
			cnt++
		} else {
			search(adj[i], end, step)
		}
	}
}
