package leetcode

/*You're given strings J representing the types of stones that are jewels, 
and S representing the stones you have.  Each character in S is a type of stone you have. 
 You want to know how many of the stones you have are also jewels.
The letters in J are guaranteed distinct, and all characters in J and S are letters. 
Letters are case sensitive, so "a" is considered a different type of stone from "A".

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/jewels-and-stones
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func numJewelsInStones(J string, S string) int {
    rlt := 0
    lenJ := len(J)
    lenS := len(S)
    for i := 0; i < lenJ; i++ {
        for j := 0; j < lenS; j++ {
            if J[i] == S[j] {
                rlt++
            }
        }
    }
    return rlt
}