package leetcode

/*Given an integer array with even length,
where different numbers in this array represent different kinds of candies.
Each number means one candy of the corresponding kind.
You need to distribute these candies equally in number to brother and sister.
Return the maximum number of kinds of candies the sister could gain.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/distribute-candies
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func distributeCandies(candies []int) int {
	mapCandy := make(map[int]int, 0)
	length := len(candies)
	rlt := 0
	for i := 0; i < length; i++ {
		_, ok := mapCandy[candies[i]]
		if !ok {
			rlt++
			mapCandy[candies[i]] = 1
		}
	}
	if rlt > length/2 {
		rlt = length / 2
	}
	return rlt
}
