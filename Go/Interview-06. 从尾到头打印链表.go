package leetcode

/*输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reversePrint(head *ListNode) []int {
    rlt := make([]int, 0, 10000)
    for head != nil {
        rlt = append(rlt, head.Val)
        head = head.Next
    }
    
    length := len(rlt) / 2
    fulLen := len(rlt) - 1
    for i := 0; i < length; i++ {
        temp := rlt[i]
        rlt[i] = rlt[fulLen - i]
        rlt[fulLen - i] = temp
    }
    return rlt
}