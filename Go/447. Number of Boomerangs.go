package leetcode

/*Given n points in the plane that are all pairwise distinct, a "boomerang" is a tuple of points (i, j, k) such that the distance between i and j equals the distance between i and k (the order of the tuple matters).

Find the number of boomerangs. You may assume that n will be at most 500 and coordinates of points are all in the range [-10000, 10000] (inclusive).*/

func numberOfBoomerangs(points [][]int) int {
	res := 0
	size := len(points)
	for i := 0; i < size; i++ {
		distance := make(map[int]int, 0)
		for j := 0; j < size; j++ {
			sqrt := (points[i][0]-points[j][0])*(points[i][0]-points[j][0]) + (points[i][1]-points[j][1])*(points[i][1]-points[j][1])
			val, ok := distance[sqrt]
			if ok {
				distance[sqrt] = val + 1
			} else {
				distance[sqrt] = 1
			}
		}
		for _, v := range distance {
			if v > 1 {
				res += v * (v - 1)
			}
		}
	}
	return res
}
