package leetcode

/*Given an integer array arr. You have to sort the integers in the array in ascending order by the number of 1's in their binary representation and in case of two or more integers have the same number of 1's you have to sort them in ascending order.

Return the sorted array.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-integers-by-the-number-of-1-bits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "sort"

type Ints []int

func sortByBits(arr []int) []int {
	var Arr Ints
	Arr = arr
	sort.Sort(Arr)
	return arr
}

func (this Ints) Len() int {
	return len(this)
}

func (this Ints) Less(i, j int) bool {
	bitsI, bitsJ := countBits(this[i]), countBits(this[j])
	if bitsI < bitsJ {
		return true
	} else if bitsI > bitsJ {
		return false
	} else {
		return this[i] < this[j]
	}
}

func (this Ints) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func countBits(n int) int {
	rlt := 0
	for n != 0 {
		if n&0x1 == 1 {
			rlt++
		}
		n = n >> 1
	}
	return rlt
}
