package leetcode

/*Alice and Bob take turns playing a game, with Alice starting first.
Initially, there is a number N on the chalkboard. 
 On each player's turn, that player makes a move consisting of:
Choosing any x with 0 < x < N and N % x == 0.
Replacing the number N on the chalkboard with N - x.
Also, if a player cannot make a move, they lose the game.
Return True if and only if Alice wins the game, assuming both players play optimally.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/divisor-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func divisorGame(N int) bool {
    return N % 2 == 0
}