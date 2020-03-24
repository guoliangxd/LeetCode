package leetcode

/*Assume you are an awesome parent and want to give your children some cookies. But, you should give each child at most one cookie. Each child i has a greed factor gi, which is the minimum size of a cookie that the child will be content with; and each cookie j has a size sj. If sj >= gi, we can assign the cookie j to the child i, and the child i will be content. Your goal is to maximize the number of your content children and output the maximum number.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/assign-cookies
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "sort"

func findContentChildren(g []int, s []int) int {
	res := 0
	leftCookie := 0
	sort.Ints(g)
	sort.Ints(s)
	for i := 0; i < len(g) && leftCookie < len(s); i++ {
		for j := leftCookie; j < len(s); j++ {
			if s[j] >= g[i] {
				res++
				leftCookie = j + 1
				break
			}
		}
	}
	return res
}
