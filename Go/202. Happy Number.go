package leetcode

/*Write an algorithm to determine if a number is "happy".

A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.*/

//快慢指针破循环
func isHappy(n int) bool {
	fast, slow := getSquareSum(n), n
	for fast != slow {
		fast = getSquareSum(fast)
		fast = getSquareSum(fast)
		slow = getSquareSum(slow)
	}
	return fast == 1
}

func getSquareSum(n int) int {
	res := 0
	for n != 0 {
		temp := n % 10
		res += temp * temp
		n /= 10
	}
	return res
}
