package leetcode

/*
Given a group of two strings, you need to find the longest uncommon subsequence of this group of two strings. The longest uncommon subsequence is defined as the longest subsequence of one of these strings and this subsequence should not be any subsequence of the other strings.

A subsequence is a sequence that can be derived from one sequence by deleting some characters without changing the order of the remaining elements. Trivially, any string is a subsequence of itself and an empty string is a subsequence of any string.

The input will be two strings, and the output needs to be the length of the longest uncommon subsequence. If the longest uncommon subsequence doesn't exist, return -1.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-uncommon-subsequence-i
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findLUSlength(a string, b string) int {
    if a == b {
        return -1
    } else if len(a) == len(b) {
        return len(a)
    } else {
        if len(a) >= len(b) {
            return len(a)
        } else {
            return len(b)
        }
    }
}