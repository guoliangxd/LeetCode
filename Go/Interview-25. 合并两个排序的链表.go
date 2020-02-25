package leetcode

/*输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    temp := &ListNode{0, nil}
    rlt := temp
    for l1 != nil && l2 != nil {
        if (l1.Val <= l2.Val) {
            temp.Next = l1
            l1 = l1.Next
        } else {
            temp.Next = l2
            l2 = l2.Next
        }
        temp = temp.Next
    }

    if l1 != nil {
        temp.Next = l1
    }

    if l2 != nil {
        temp.Next = l2
    }

    return rlt.Next
}