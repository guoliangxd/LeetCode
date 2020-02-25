package leetcode

/*定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    var pre *ListNode
    for head != nil {
        temp := head
        head = head.Next

        temp.Next = pre
        pre = temp  
    }
    return pre
}