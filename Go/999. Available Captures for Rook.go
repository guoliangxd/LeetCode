package leetcode

/*On an 8 x 8 chessboard, there is one white rook.  There also may be empty squares, white bishops, and black pawns.  These are given as characters 'R', '.', 'B', and 'p' respectively. Uppercase characters represent white pieces, and lowercase characters represent black pieces.

The rook moves as in the rules of Chess: it chooses one of four cardinal directions (north, east, west, and south), then moves in that direction until it chooses to stop, reaches the edge of the board, or captures an opposite colored pawn by moving to the same square it occupies.  Also, rooks cannot move into the same square as other friendly bishops.

Return the number of pawns the rook can capture in one move.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/available-captures-for-rook
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func numRookCaptures(board [][]byte) int {
	res := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 'R' {
				for k := i - 1; k >= 0; k-- {
					if board[k][j] == 'p' {
						res++
						break
					}
					if board[k][j] == 'B' {
						break
					}
				}
				for k := i + 1; k < 8; k++ {
					if board[k][j] == 'p' {
						res++
						break
					}
					if board[k][j] == 'B' {
						break
					}
				}
				for k := j - 1; k >= 0; k-- {
					if board[i][k] == 'p' {
						res++
						break
					}
					if board[i][k] == 'B' {
						break
					}
				}
				for k := j + 1; k < 8; k++ {
					if board[i][k] == 'p' {
						res++
						break
					}
					if board[i][k] == 'B' {
						break
					}
				}
				return res
			}
		}
	}
	return res
}
