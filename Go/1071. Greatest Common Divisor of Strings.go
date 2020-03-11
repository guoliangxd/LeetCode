package leetcode

/*For strings S and T, we say "T divides S" if and only if S = T + ... + T  (T concatenated with itself 1 or more times)

Return the largest string X such that X divides str1 and X divides str2.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/greatest-common-divisor-of-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func gcdOfStrings(str1 string, str2 string) string {
    lgcd := getGCD(len(str1), len(str2))
    if str1[ : lgcd] == str2[ : lgcd] && isGCD(str1, str1[ : lgcd]) && isGCD(str2, str2[ : lgcd]) {
        return str1[ : lgcd]
    } else {
        return ""
    }
}

func getGCD(m, n int) int {
    if m == n {
        return m
    } else if m > n {
        return getGCD(n, m - n)
    } else {
        return getGCD(m, n - m)
    }
}

func isGCD(str, gcd string) bool {
    lstr := len(str)
    lgcd := len(gcd)
    cycle := lstr / lgcd
    for i := 0; i < cycle; i++ {
        if str[i * lgcd : (i + 1) * lgcd] != gcd {
            return false 
        }
    }
    return true
}