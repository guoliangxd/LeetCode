package leetcode

/*You need to construct a string consists of parenthesis and integers from a binary tree with the preorder traversing way.

The null node needs to be represented by empty parenthesis pair "()". And you need to omit all the empty parenthesis pairs that don't affect the one-to-one mapping relationship between the string and the original binary tree.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/construct-string-from-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}

	this := strconv.Itoa(t.Val)
	if t.Left == nil && t.Right == nil {
		return this
	} else if t.Left != nil && t.Right != nil {
		return this + "(" + tree2str(t.Left) + ")" + "(" + tree2str(t.Right) + ")"
	} else if t.Left != nil && t.Right == nil {
		return this + "(" + tree2str(t.Left) + ")"
	} else {
		return this + "()" + "(" + tree2str(t.Right) + ")"
	}
}
