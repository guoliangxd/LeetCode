package leetcode

/*输入一棵二叉树的根节点，判断该树是不是平衡二叉树。
如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 import "math
 var isBalance bool
 func isBalanced(root *TreeNode) bool {
	 isBalance = true
	 getDepth(root)
	 return isBalance
 }
 
 func getDepth(root *TreeNode) int {
	 if !isBalance {
		 return 0
	 }
 
	 if root == nil {
		 return 0
	 }
 
	 Left := getDepth(root.Left)
	 Right := getDepth(root.Right)
	 if Left - Right > 1 || Right - Left > 1 {
		 isBalance = false
	 }
	 return int(math.Max(float64(Right), float64(Left))) + 1
 }