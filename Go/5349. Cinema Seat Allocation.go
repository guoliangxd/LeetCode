package leetcode

/*A cinema has n rows of seats, numbered from 1 to n and there are ten seats in each row, labelled from 1 to 10 as shown in the figure above.

Given the array reservedSeats containing the numbers of seats already reserved, for example, reservedSeats[i]=[3,8] means the seat located in row 3 and labelled with 8 is already reserved.

Return the maximum number of four-person families you can allocate on the cinema seats. A four-person family occupies fours seats in one row, that are next to each other. Seats across an aisle (such as [3,3] and [3,4]) are not considered to be next to each other, however, It is permissible for the four-person family to be separated by an aisle, but in that case, exactly two people have to sit on each side of the aisle.*/

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	reserved := make(map[int][]int, 0)
	res := 0
	for i := 0; i < len(reservedSeats); i++ {
		row, ok := reserved[reservedSeats[i][0]]
		if !ok {
			data := make([]int, 10)
			data[reservedSeats[i][1]-1] = 1
			reserved[reservedSeats[i][0]] = data
		} else {
			row[reservedSeats[i][1]-1] = 1
		}
	}

	res += (n - len(reserved)) * 2
	for _, row := range reserved {
		left, mid := false, false
		if row[1] == 0 && row[2] == 0 && row[3] == 0 && row[4] == 0 {
			res++
			left = true
		}

		if !left && row[3] == 0 && row[4] == 0 && row[5] == 0 && row[6] == 0 {
			res++
			mid = true
		}

		if !mid && row[5] == 0 && row[6] == 0 && row[7] == 0 && row[8] == 0 {
			res++
		}
	}

	return res
}
