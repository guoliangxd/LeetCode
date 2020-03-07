package leetcode

/*There are some chips, and the i-th chip is at position chips[i].

You can perform any of the two following types of moves any number of times (possibly zero) on any chip:

Move the i-th chip by 2 units to the left or to the right with a cost of 0.
Move the i-th chip by 1 unit to the left or to the right with a cost of 1.
There can be two or more chips at the same position initially.

Return the minimum cost needed to move all the chips to the same position (any position).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/play-with-chips
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func minCostToMoveChips(chips []int) int {
    odd, even, cost := 0, 0, 0
    for i := 0; i < len(chips); i++ {
        if chips[i] % 2 == 0 {
            even++
        } else {
            odd++
        }
    }
    if even > odd {
        cost = odd
    } else {
        cost = even
    } 

    return cost
}