package leetcode

/*Given an integer n. Each number from 1 to n is grouped according to the sum of its digits.

Return how many groups have the largest size.*/

func countLargestGroup(n int) int {
	cnt := make(map[int]int, 0)
	max := 0
	for i := 1; i <= n; i++ {
		temp := count(i)
		v := cnt[temp]
		v++
		cnt[temp] = v

		if v > max {
			max = v
		}
	}

	res := 0
	for _, v := range cnt {
		if v == max {
			res++
		}
	}
	return res
}

func count(n int) int {
	res := 0
	for n != 0 {
		res += n % 10
		n /= 10
	}
	return res
}
