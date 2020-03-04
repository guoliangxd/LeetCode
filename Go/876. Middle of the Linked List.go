package leetcode

/*Given a non-empty, singly linked list with head node head, return a middle node of linked list.

If there are two middle nodes, return the second middle node.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/middle-of-the-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	temp := &ListNode{0, head}
	slow, fast := temp, temp
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		fast = fast.Next

		slow = slow.Next
	}
	if fast == nil {
		return slow
	} else {
		return slow.Next
	}

}
