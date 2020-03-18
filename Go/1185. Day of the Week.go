package leetcode

/*Given a date, return the corresponding day of the week for that date.

The input is given as three integers representing the day, month and year respectively.

Return the answer as one of the following values {"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}.*/

func dayOfTheWeek(day int, month int, year int) string {
	months := [...]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	days := day
	for i := 1971; i < year; i++ {
		if isLeap(i) {
			days += 366
		} else {
			days += 365
		}
	}

	for i := 0; i < month-1; i++ {
		days += months[i]
	}

	if month > 2 && isLeap(year) {
		days++
	}

	res := ""
	switch days % 7 {
	case 1:
		res = "Friday"
	case 2:
		res = "Saturday"
	case 3:
		res = "Sunday"
	case 4:
		res = "Monday"
	case 5:
		res = "Tuesday"
	case 6:
		res = "Wednesday"
	case 0:
		res = "Thursday"
	}
	return res
}

func isLeap(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	} else {
		return year%4 == 0
	}
}
