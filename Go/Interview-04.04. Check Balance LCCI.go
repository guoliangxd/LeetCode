package leetcode

/*Implement a function to check if a binary tree is balanced. For the purposes of this question, a balanced tree is defined to be a tree such that the heights of the two subtrees of any node never differ by more than one.
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/check-balance-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

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