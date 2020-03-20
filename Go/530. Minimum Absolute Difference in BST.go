package leetcode

/*Given a binary search tree with non-negative values, find the minimum absolute difference between values of any two nodes.*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var min int
var last int

func getMinimumDifference(root *TreeNode) int {
	last = -0x7FFFFFFF - 1
	min = 0x7FFFFFFF
	inOrder(root)
	return min
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}

	inOrder(root.Left)

	if root.Val-last < min {
		min = root.Val - last
	}
	last = root.Val
	inOrder(root.Right)
}
