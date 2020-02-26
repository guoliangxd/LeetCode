package leetcode

/*The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
Given two integers x and y, calculate the Hamming distance.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/hamming-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func hammingDistance(x int, y int) int {
    rlt := 0
    /*
    xor := x ^ y
    for xor != 0 {
        if xor % 2 == 1 {
            rlt++
        }
        xor = xor >> 1
    }*/

    for x != 0 || y != 0 {
        if (x % 2) ^ (y % 2) == 1 {
            rlt++
        }
        x = x >> 1
        y = y >> 1
    }
    return rlt
}