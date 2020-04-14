package leetcode

/*You are given two non-empty linked lists representing two non-negative integers. The most significant digit comes first and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Follow up:
What if you cannot modify the input lists? In other words, reversing the lists is not allowed.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var res *ListNode
    arr1 := make([]int, 0)
    arr2 := make([]int, 0)
    if l1.Val == 0 {
        return l2 
    }
    if l2.Val == 0 {
        return l1
    }
    for l1 != nil {
        arr1 = append(arr1, l1.Val)
        l1 = l1.Next
    }
    for l2 != nil {
        arr2 = append(arr2, l2.Val)
        l2 = l2.Next
    }
    size1 := len(arr1)
    size2 := len(arr2)
    carry := 0
    i := 0
    for ; i < size1 && i < size2; i++ {
        value := arr1[size1 - i -1] + arr2[size2 - i - 1] + carry
        value, carry = value % 10, value / 10
        newNode := &ListNode{value, res}
        res = newNode
    }
    for ; i < size1; i++ {
        value := arr1[size1 - i - 1] + carry
        value, carry = value % 10, value / 10
        newNode := &ListNode{value, res}
        res = newNode
    }
    for ; i < size2; i++ {
        value := arr2[size2 - i - 1] + carry
        value, carry = value % 10, value / 10
        newNode := &ListNode{value, res}
        res = newNode
    }
    if carry != 0 {
        newNode := &ListNode{carry, res}
        res = newNode
    }
    return res
}

