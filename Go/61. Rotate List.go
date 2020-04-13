package leetcode

/*Given a linked list, rotate the list to the right by k places, where k is non-negative.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil || k == 0 {
        return head
    }

    size := 1
    cur := head
    //循环结束后cur指向当前链表的尾节点
    for cur.Next != nil {
        size++
        cur = cur.Next
    }

    if offset := k % size; offset == 0 {
        return head
    } else {
        temp := head
        for i := 0; i < size - offset - 1; i++ {
            temp = temp.Next
        }
        res := temp.Next
        temp.Next = nil
        cur.Next = head
        return res
    }
}

