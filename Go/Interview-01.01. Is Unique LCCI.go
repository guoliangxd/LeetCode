package leetcode

/*Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/is-unique-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

//使用了sort包的Sort方法
import "sort"

type String []rune

func isUnique(astr string) bool {
	temp := String(astr)
	sort.Sort(temp)
	length := len(temp)
	for i := 0; i < length-1; i++ {
		if temp[i] == temp[i+1] {
			return false
		}
	}
	return true
}

func (this String) Len() int {
	return len(this)
}

func (this String) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this String) Swap(i, j int) {
	temp := this[i]
	this[i] = this[j]
	this[j] = temp
}
