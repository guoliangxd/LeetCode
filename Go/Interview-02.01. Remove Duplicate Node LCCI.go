package leetcode

/*Write code to remove duplicates from an unsorted linked list.*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeDuplicateNodes(head *ListNode) *ListNode {
	mapList := make(map[int]int, 0)
	cur, prv := head, head
	for cur != nil {
		_, ok := mapList[cur.Val]
		if ok {
			prv.Next = cur.Next
		} else {
			mapList[cur.Val] = 1
			prv = cur
		}
		cur = cur.Next
	}
	return head
}
