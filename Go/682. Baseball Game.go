package leetcode

/*You're now a baseball game point recorder.

Given a list of strings, each string can be one of the 4 following types:

Integer (one round's score): Directly represents the number of points you get in this round.
"+" (one round's score): Represents that the points you get in this round are the sum of the last two valid round's points.
"D" (one round's score): Represents that the points you get in this round are the doubled data of the last valid round's points.
"C" (an operation, which isn't a round's score): Represents the last valid round's points you get were invalid and should be removed.
Each round's operation is permanent and could have an impact on the round before and the round after.

You need to return the sum of the points you could get in all the rounds.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/baseball-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "strconv"
func calPoints(ops []string) int {
    rlt := 0
    score := make([]int, 0, 1)
    for _, v := range ops {
        switch v {
            case "+":
                score = append(score, score[len(score) - 2] + score[len(score) - 1])
            case "D":
                score = append(score, score[len(score) - 1] * 2)
            case "C":
                score = score[ : len(score) - 1]
            default:
                val, _ := strconv.Atoi(v)
                score = append(score, val)       
        }
    }

    for _, v := range score {
        rlt += v
    }
    return rlt
}