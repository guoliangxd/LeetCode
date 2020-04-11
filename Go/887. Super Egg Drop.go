package leetcode

import "math"

/*You are given K eggs, and you have access to a building with N floors from 1 to N.

Each egg is identical in function, and if an egg breaks, you cannot drop it again.

You know that there exists a floor F with 0 <= F <= N such that any egg dropped at a floor higher than F will break, and any egg dropped at or below floor F will not break.

Each move, you may take an egg (if you have an unbroken one) and drop it from any floor X (with 1 <= X <= N).

Your goal is to know with certainty what the value of F is.

What is the minimum number of moves that you need to know with certainty what F is, regardless of the initial value of F?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/super-egg-drop
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

var memo map[int]int

func superEggDrop(K int, N int) int {
	memo = make(map[int]int)
	return dp(K, N)
}

func dp(K, N int) int {
	if _, ok := memo[N*100+K]; !ok {
		ans := 0
		if N == 0 {
			ans = 0
		} else if K == 1 {
			ans = N
		} else {
			lo, hi := 1, N
			for lo+1 < hi {
				x := (lo + hi) / 2
				t1 := dp(K-1, x-1)
				t2 := dp(K, N-x)

				if t1 < t2 {
					lo = x
				} else if t1 > t2 {
					hi = x
				} else {
					lo, hi = x, x
				}
			}
			ans = 1 + int(math.Min(math.Max(float64(dp(K-1, lo-1)), float64(dp(K, N-lo))), math.Max(float64(dp(K-1, hi-1)), float64(dp(K, N-hi)))))
		}
		memo[N*100+K] = ans
	}
	return memo[N*100+K]
}
