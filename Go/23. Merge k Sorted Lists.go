package leetcode

/*Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	res := &ListNode{0, nil}
	head := res
	MAX := 0x7fffffff
	for {
		find := false
		cur := MAX
		for i := 0; i < len(lists); i++ {
			if lists[i] != nil {
				find = true
				if lists[i].Val < cur {
					cur = lists[i].Val
				}
			} else {
				// 删除空链表
				lists = append(lists[:i], lists[i+1:]...)
				i--
			}
		}
		if !find {
			// 链表为空则退出循环
			break
		}
		for i := 0; i < len(lists); i++ {
			for lists[i] != nil && lists[i].Val == cur {
				res.Next = lists[i]
				lists[i] = lists[i].Next
				res = res.Next
				res.Next = nil
			}
		}
	}

	return head.Next
}
