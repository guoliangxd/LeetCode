package leetcode

/*Given a string date representing a Gregorian calendar date formatted as YYYY-MM-DD, return the day number of the year.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/day-of-the-year
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import (
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	res := 0
	months := [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	dates := strings.Split(date, "-")
	year, _ := strconv.Atoi(dates[0])
	month, _ := strconv.Atoi(dates[1])
	day, _ := strconv.Atoi(dates[2])

	res += day
	for i := 0; i < month-1; i++ {
		res += months[i]
	}

	if month > 2 {
		if year%100 == 0 {
			if year%400 == 0 {
				res++
			}
		} else if year%4 == 0 {
			res++
		}
	}
	return res
}
