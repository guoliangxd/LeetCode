package leetcode

/*The data structure TreeNode is used for binary tree, but it can also used to represent a single linked list (where left is null, and right is the next node in the list). Implement a method to convert a binary search tree (implemented with TreeNode) into a single linked list. The values should be kept in order and the operation should be performed in place (that is, on the original data structure).

Return the head node of the linked list after converting.

Note: This problem is slightly different from the original one in the book.*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res, cur *TreeNode

func convertBiNode(root *TreeNode) *TreeNode {
	res = &TreeNode{0, nil, nil}
	cur = res
	//中序遍历二叉树
	inOrder(root)
	return res.Right
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}

	inOrder(root.Left)
	cur.Right = root
	cur = root
	root.Left = nil
	inOrder(root.Right)
}
