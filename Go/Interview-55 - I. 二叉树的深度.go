package leetcode

/*输入一棵二叉树的根节点，求该树的深度。
从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		leftLen := maxDepth(root.Left)
		rightLen := maxDepth(root.Right)
		if leftLen >= rightLen {
			return leftLen + 1
		} else {
			return rightLen + 1
		}
	}
}
