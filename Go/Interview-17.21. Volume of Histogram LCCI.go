package leetcode

/*Imagine a histogram (bar graph). Design an algorithm to compute the volume of water it could hold if someone poured water across the top. You can assume that each histogram bar has width 1.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/volume-of-histogram-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func trap(height []int) int {
	size := len(height)
	lMax := make([]int, size)
	rMax := make([]int, size)
	lmax := 0
	rmax := 0

	for i := 0; i < size; i++ {
		if height[i] > lmax {
			lmax = height[i]
			lMax[i] = height[i]
		} else {
			lMax[i] = lmax
		}

		j := size - i - 1
		if height[j] > rmax {
			rmax = height[j]
			rMax[j] = height[j]
		} else {
			rMax[j] = rmax
		}
	}

	res := 0

	for i := 0; i < size; i++ {
		min := lMax[i]
		if rMax[i] < min {
			min = rMax[i]
		}
		res += min - height[i]
	}

	return res
}
