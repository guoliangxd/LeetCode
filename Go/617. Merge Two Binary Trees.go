package leetcode

/*Given two binary trees and imagine that when you put one of them to cover the other,
some nodes of the two trees are overlapped while the others are not.

You need to merge them into a new binary tree.
The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node.
 Otherwise, the NOT null node will be used as the node of new tree.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-binary-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	} else if t1 != nil && t2 == nil {
		t1.Left = mergeTrees(t1.Left, nil)
		t1.Right = mergeTrees(t1.Right, nil)
		return t1
	} else if t1 == nil && t2 != nil {
		t2.Left = mergeTrees(nil, t2.Left)
		t2.Right = mergeTrees(nil, t2.Right)
		return t2
	} else {
		t1.Val += t2.Val
		t1.Left = mergeTrees(t1.Left, t2.Left)
		t1.Right = mergeTrees(t1.Right, t2.Right)
		return t1
	}
}
