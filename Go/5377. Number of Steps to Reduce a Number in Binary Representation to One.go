package leetcode

/*Given a number s in their binary representation. Return the number of steps to reduce it to 1 under the following rules:

If the current number is even, you have to divide it by 2.

If the current number is odd, you have to add 1 to it.

It's guaranteed that you can always reach to one for all testcases.*/

func numSteps(s string) int {
    num := []byte(s)
    num = append([]byte{0}, num...)
    
    bits := 0
    for i := 0; i < len(num); i++ {
        if num[i] == '1' {
            bits++
        }
    }
    
    cnt := 0
    for !(bits == 1 && num[len(num) - 1] == '1') {
        if num[len(num) - 1] == '0' {
            move(num)
        } else {
            for i := len(num) - 1; i >= 0; i-- {
                if num[i] == '1' {
                    num[i] = '0'
                    bits--
                } else {
                    num[i] = '1'
                    bits++
                    break
                }
            }
        }
        cnt++
    }
    return cnt
}

func move(data []byte) {
    for i := len(data) - 1; i > 0; i-- {
        data[i] = data[i - 1]
    }
    data[0] = '0'
}