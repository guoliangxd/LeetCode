package leetcode

/*You have an array of logs.  Each log is a space delimited string of words.

For each log, the first word in each log is an alphanumeric identifier.  Then, either:

Each word after the identifier will consist only of lowercase letters, or;
Each word after the identifier will consist only of digits.
We will call these two varieties of logs letter-logs and digit-logs.  It is guaranteed that each log has at least one word after its identifier.

Reorder the logs so that all of the letter-logs come before any digit-log.  The letter-logs are ordered lexicographically ignoring identifier, with the identifier used in case of ties.  The digit-logs should be put in their original order.

Return the final order of the logs.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reorder-data-in-log-files
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

type Data []string

func reorderLogFiles(logs []string) []string {
	BubleSort(Data(logs))
	return logs
}

func (this Data) Less(i, j int) bool {
	iSpace := 0
	jSpace := 0
	for this[i][iSpace] != ' ' {
		iSpace++
	}
	for this[j][jSpace] != ' ' {
		jSpace++
	}
	iSpace++
	jSpace++

	iAlpah := this[i][iSpace] >= 'a' && this[i][iSpace] <= 'z'
	jAlpha := this[j][jSpace] >= 'a' && this[j][jSpace] <= 'z'

	if iAlpah && jAlpha {
		res := cmpStr(this[i][iSpace:], this[j][jSpace:])
		if res == -1 {
			return true
		} else if res == 1 {
			return false
		}

		res = cmpStr(this[i][:iSpace], this[j][:jSpace])
		if res == -1 {
			return true
		} else {
			return false
		}
	} else if !iAlpah && !jAlpha {
		return true
	} else if iAlpah && !jAlpha {
		return true
	} else {
		return false
	}

}

func (this Data) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this Data) Len() int {
	return len(this)
}

func cmpStr(s1, s2 string) int {
	size1 := len(s1)
	size2 := len(s2)
	size := size1
	if size2 < size {
		size = size2
	}
	for i := 0; i < size; i++ {
		if s1[i] < s2[i] {
			return -1
		} else if s1[i] > s2[i] {
			return 1
		}
	}

	if size1 == size2 {
		return 0
	} else if size1 < size2 {
		return -1
	} else {
		return 1
	}
}

func BubleSort(data Data) {
	size := len(data)
	for i := 0; i < size-1; i++ {
		swaped := false
		for j := 0; j < size-i-1; j++ {
			if !data.Less(j, j+1) {
				data.Swap(j, j+1)
				swaped = true
			}
		}
		if !swaped {
			break
		}
	}
}
