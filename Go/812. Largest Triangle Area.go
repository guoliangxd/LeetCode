package leetcode

import "math"

/*You have a list of points in the plane. Return the area of the largest triangle that can be formed by any 3 of the points.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-triangle-area
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func largestTriangleArea(points [][]int) float64 {
	n := len(points)
	res := 0.0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				res = math.Max(res, area(points[i], points[j], points[k]))
			}
		}
	}
	return res
}

//用鞋带公式求简单多边形面积
func area(P, Q, R []int) float64 {
	return 0.5 * math.Abs(float64(P[0]*Q[1]+Q[0]*R[1]+R[0]*P[1]-P[1]*Q[0]-Q[1]*R[0]-R[1]*P[0]))
}
