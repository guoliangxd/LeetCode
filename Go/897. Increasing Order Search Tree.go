package leetcode

/*Given a binary search tree, rearrange the tree in in-order
so that the leftmost node in the tree is now the root of the tree,
and every node has no left child and only 1 right child.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/increasing-order-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var cur *TreeNode

func increasingBST(root *TreeNode) *TreeNode {
	cur = &TreeNode{0, nil, nil}
	head := cur
	inOrder(root)
	return head.Right
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	inOrder(root.Left)
	cur.Right = root
	root.Left = nil
	cur = root
	inOrder(root.Right)
}
