package leetcode

/*Reverse a singly linked list.

Example:
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//迭代解法
 func reverseList(head *ListNode) *ListNode {
    cur := head
    var prv *ListNode
    for cur != nil {
        temp := cur.Next
        cur.Next = prv
        prv = cur
        cur = temp
    }
    return prv
}