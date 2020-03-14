package leetcode

/*给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。

返回删除后的链表的头节点。*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	temp := &ListNode{val ^ 0XFFFFFFFF, head}
	cur, pre := temp, temp
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
			break
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return temp.Next
}
