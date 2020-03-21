package leetcode

/*You are given two jugs with capacities x and y litres. There is an infinite amount of water supply available. You need to determine whether it is possible to measure exactly z litres using these two jugs.

If z liters of water is measurable, you must have z liters of water contained within one or both buckets by the end.

Operations allowed:

Fill any of the jugs completely with water.
Empty any of the jugs.
Pour water from one jug into another till the other jug is completely full or the first jug itself is empty.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/water-and-jug-problem
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func canMeasureWater(x int, y int, z int) bool {
	if x == 0 || y == 0 {
		if z == 0 {
			return true
		} else {
			if x > 0 {
				return z%x == 0
			} else if y > 0 {
				return z%y == 0
			} else {
				return false
			}
		}
	}
	if x+y < z {
		return false
	} else {
		return z%gcd(x, y) == 0
	}
}

func gcd(x, y int) int {
	if x%y == 0 {
		return y
	} else {
		return gcd(y, x%y)
	}
}
