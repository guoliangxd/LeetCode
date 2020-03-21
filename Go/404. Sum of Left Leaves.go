package leetcode

/*Find the sum of all left leaves in a given binary tree.*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	if root == nil {
		return res
	} else {
		res += sumLeft(root.Left, true)
		res += sumLeft(root.Right, false)
	}
	return res
}

func sumLeft(root *TreeNode, isLeft bool) int {
	if root == nil {
		return 0
	}

	if isLeft && root.Left == nil && root.Right == nil {
		return root.Val
	}

	return sumLeft(root.Left, true) + sumLeft(root.Right, false)
}
