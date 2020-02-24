package leetcode

import "math"

/*On a plane there are n points with integer coordinates points[i] = [xi, yi].
 Your task is to find the minimum time in seconds to visit all points.
You can move according to the next rules:
In one second always you can either move vertically,
horizontally by one unit or diagonally (it means to move one unit vertically and one unit horizontally in one second).
You have to visit the points in the same order as they appear in the array.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-time-visiting-all-points
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func minTimeToVisitAllPoints(points [][]int) int {
	//使用math包比不使用math包效率要高很多
	rlt := 0
	length := len(points) - 1
	for i := 0; i < length; i++ {
		/*
		   difX := points[i][0] - points[i + 1][0]
		   if difX < 0 {
		       difX = -difX
		   }
		   difY := points[i][1] - points[i + 1][1]
		   if difY < 0 {
		       difY = -difY
		   }
		   if difX >= difY {
		       rlt += difX
		   } else {
		       rlt += difY
		   }
		*/
		rlt += int(math.Max(math.Abs(float64(points[i][0]-points[i+1][0])), math.Abs(float64(points[i][1]-points[i+1][1]))))
	}
	return rlt
}
