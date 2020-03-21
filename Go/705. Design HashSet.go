package leetcode

/*Design a HashSet without using any built-in hash table libraries.

To be specific, your design should include these functions:

add(value): Insert a value into the HashSet.
contains(value) : Return whether the value exists in the HashSet or not.
remove(value): Remove a value in the HashSet. If the value does not exist in the HashSet, do nothing.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/design-hashset
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

type Node struct {
	Val  int
	Next *Node
}

type MyHashSet struct {
	data [1111111]*Node
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		[1111111]*Node{},
	}
}

func (this *MyHashSet) Add(key int) {
	if this.Contains(key) {
		return
	}

	hashCode := key % 1111111
	this.data[hashCode] = &Node{
		key,
		this.data[hashCode],
	}
}

func (this *MyHashSet) Remove(key int) {
	hashCode := key % 1111111
	cur := this.data[hashCode]
	temp := &Node{
		0,
		cur,
	}
	prv := temp
	for cur != nil {
		if cur.Val == key {
			prv.Next = cur.Next
			break
		}
		cur = cur.Next
		prv = prv.Next
	}
	this.data[hashCode] = temp.Next
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	bucket := this.data[key%1111111]
	for bucket != nil {
		if bucket.Val == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
