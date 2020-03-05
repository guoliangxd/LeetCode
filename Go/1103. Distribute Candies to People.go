package leetcode

/*We distribute some number of candies, to a row of n = num_people people in the following way:

We then give 1 candy to the first person, 2 candies to the second person,
and so on until we give n candies to the last person.

Then, we go back to the start of the row, giving n + 1 candies to the first person,
n + 2 candies to the second person, and so on until we give 2 * n candies to the last person.

This process repeats (with us giving one more candy each time,
and moving to the start of the row after we reach the end) until we run out of candies.
The last person will receive all of our remaining candies
(not necessarily one more than the previous gift).

Return an array (of length num_people and sum candies)
that represents the final distribution of candies.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/distribute-candies-to-people
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func distributeCandies(candies int, num_people int) []int {
	rlt := make([]int, num_people, num_people)
	cnt := 0
	for candies > 0 {
		temp := 0
		if cnt+1 <= candies {
			temp = cnt + 1
		} else {
			temp = candies
		}
		rlt[cnt%num_people] += temp
		candies -= temp
		cnt++
	}
	return rlt
}
