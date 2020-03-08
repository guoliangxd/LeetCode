package leetcode

/*A binary tree is univalued if every node in the tree has the same value.

Return true if and only if the given tree is univalued.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/univalued-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	if root.Left != nil {
		if root.Left.Val != root.Val {
			return false
		}
	}

	if root.Right != nil {
		if root.Right.Val != root.Val {
			return false
		}
	}

	return isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
