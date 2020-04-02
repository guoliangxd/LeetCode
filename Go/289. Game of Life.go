package leetcode

/*According to the Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."

Given a board with m by n cells, each cell has an initial state live (1) or dead (0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

Any live cell with fewer than two live neighbors dies, as if caused by under-population.
Any live cell with two or three live neighbors lives on to the next generation.
Any live cell with more than three live neighbors dies, as if by over-population..
Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
Write a function to compute the next state (after one update) of the board given its current state. The next state is created by applying the above rules simultaneously to every cell in the current state, where births and deaths occur simultaneously.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/game-of-life
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func gameOfLife(board [][]int) {
	change := make([][]int, 0)
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			alive := 0
			if i-1 >= 0 {
				if board[i-1][j] == 1 {
					alive++
				}
				if j-1 >= 0 && board[i-1][j-1] == 1 {
					alive++
				}
				if j+1 < n && board[i-1][j+1] == 1 {
					alive++
				}
			}
			if i+1 < m {
				if board[i+1][j] == 1 {
					alive++
				}
				if j-1 >= 0 && board[i+1][j-1] == 1 {
					alive++
				}
				if j+1 < n && board[i+1][j+1] == 1 {
					alive++
				}
			}
			if j-1 >= 0 && board[i][j-1] == 1 {
				alive++
			}
			if j+1 < n && board[i][j+1] == 1 {
				alive++
			}
			if board[i][j] == 1 && (alive < 2 || alive > 3) {
				change = append(change, []int{i, j})
			}
			if board[i][j] == 0 && alive == 3 {
				change = append(change, []int{i, j})
			}
		}
	}

	for i := 0; i < len(change); i++ {
		board[change[i][0]][change[i][1]] ^= 1
	}
}
