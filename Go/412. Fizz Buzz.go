package leetcode

/*Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”. For numbers which are multiples of both three and five output “FizzBuzz”.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/fizz-buzz
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "strconv"
func fizzBuzz(n int) []string {
    res := make([]string, n)

    three, five := 1, 1
    for i := 0; i < n; i++ {
        if three == 3 && five == 5 {
            res[i] = "FizzBuzz"
            three = 0
            five = 0
        } else if three == 3 {
            three = 0
            res[i] = "Fizz"
        } else if five == 5 {
            five = 0
            res[i] = "Buzz"
        } else {
            res[i] = strconv.Itoa(i + 1)
        }
        three++
        five++
    }
    return res
}