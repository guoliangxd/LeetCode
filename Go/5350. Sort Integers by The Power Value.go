package leetcode

/*The power of an integer x is defined as the number of steps needed to transform x into 1 using the following steps:

if x is even then x = x / 2
if x is odd then x = 3 * x + 1
For example, the power of x = 3 is 7 because 3 needs 7 steps to become 1 (3 --> 10 --> 5 --> 16 --> 8 --> 4 --> 2 --> 1).

Given three integers lo, hi and k. The task is to sort all integers in the interval [lo, hi] by the power value in ascending order, if two or more integers have the same power value sort them by ascending order.

Return the k-th integer in the range [lo, hi] sorted by the power value.

Notice that for any integer x (lo <= x <= hi) it is guaranteed that x will transform into 1 using these steps and that the power of x is will fit in 32 bit signed integer.*/

import "sort"

type data struct {
	Val    int
	Weight int
}
type Data []data

func (this Data) Len() int {
	return len(this)
}

func (this Data) Less(i, j int) bool {
	if this[i].Weight < this[j].Weight {
		return true
	} else if this[i].Weight == this[j].Weight {
		return this[i].Val < this[j].Val
	} else {
		return false
	}
}

func (this Data) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
func getKth(lo int, hi int, k int) int {
	arr := make(Data, hi-lo+1)
	for i := 0; i < len(arr); i++ {
		arr[i] = data{
			lo + i,
			weight(lo + i),
		}
	}
	sort.Sort(arr)
	return arr[k-1].Val
}

func weight(n int) int {
	res := 0
	for n != 1 {
		if n%2 == 1 {
			n = n*3 + 1
		} else {
			n = n / 2
		}
		res++
	}
	return res
}
