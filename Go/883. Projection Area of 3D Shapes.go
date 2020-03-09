package leetcode

/*On a N * N grid, we place some 1 * 1 * 1 cubes that are axis-aligned with the x, y, and z axes.

Each value v = grid[i][j] represents a tower of v cubes placed on top of grid cell (i, j).

Now we view the projection of these cubes onto the xy, yz, and zx planes.

A projection is like a shadow, that maps our 3 dimensional figure to a 2 dimensional plane.

Here, we are viewing the "shadow" when looking at the cubes from the top, the front, and the side.

Return the total area of all three projections.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/projection-area-of-3d-shapes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func projectionArea(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	x, y, z := 0, 0, 0

	for i := 0; i < m; i++ {
		maxOnX := 0
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				z++

				if maxOnX < grid[i][j] {
					maxOnX = grid[i][j]
				}
			}
		}
		x += maxOnX
	}

	for j := 0; j < n; j++ {
		maxOnY := 0
		for i := 0; i < m; i++ {
			if grid[i][j] > maxOnY {
				maxOnY = grid[i][j]
			}
		}
		y += maxOnY
	}
	return x + y + z
}
