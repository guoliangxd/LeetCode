package leetcode

/*Write a program to find the node at which the intersection of two singly linked lists begins.*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	node1, node2 := headA, headB
	for node1 != node2 {
		if node1 == nil {
			node1 = headB
		} else {
			node1 = node1.Next
		}

		if node2 == nil {
			node2 = headA
		} else {
			node2 = node2.Next
		}
	}
	return node1
}
