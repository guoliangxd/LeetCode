package leetcode

/*Given a binary tree, you need to compute the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diameter-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 var res int
 func diameterOfBinaryTree(root *TreeNode) int {
	 if root == nil {
		 return 0
	 }
 
	 res = 0
 
	 depth(root)
	 
	 return res
 }
 
 func depth(root *TreeNode) int {
	 if root == nil {
		 return 0
	 }
 
	 L := depth(root.Left)
	 R := depth(root.Right)
	 
	 if res < L + R {
		 res = L + R
	 }
	 
	 if L >= R {
		 return L + 1
	 } else {
		 return R + 1
	 }
 }