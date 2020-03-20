package leetcode

/*Design a HashMap without using any built-in hash table libraries.

To be specific, your design should include these functions:

put(key, value) : Insert a (key, value) pair into the HashMap. If the value already exists in the HashMap, update the value.
get(key): Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
remove(key) : Remove the mapping for the value key if this map contains the mapping for the key.
*/

type Node struct {
	K, V int
	Next *Node
}

type MyHashMap struct {
	buckets [1111]*Node
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{
		[1111]*Node{},
	}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	hashCode := key % 1111
	bucket := this.buckets[hashCode]
	find := false
	for bucket != nil {
		if bucket.K == key {
			bucket.V = value
			find = true
			break
		}
		bucket = bucket.Next
	}
	if !find {
		temp := this.buckets[hashCode]
		this.buckets[hashCode] = &Node{
			key,
			value,
			temp,
		}
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	hashCode := key % 1111
	bucket := this.buckets[hashCode]
	for bucket != nil {
		if bucket.K == key {
			return bucket.V
		}
		bucket = bucket.Next
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	hashCode := key % 1111
	temp := &Node{
		0,
		0,
		this.buckets[hashCode],
	}
	pre := temp
	cur := temp.Next
	for cur != nil {
		if cur.K == key {
			pre.Next = cur.Next
			break
		}
		pre, cur = cur, cur.Next
	}
	this.buckets[hashCode] = temp.Next
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
