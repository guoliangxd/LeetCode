package leetcode

/*Given a binary tree, return all root-to-leaf paths.

Note: A leaf is a node with no children.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-paths
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 import "strconv
 var res []string
 func binaryTreePaths(root *TreeNode) []string {
	 if root == nil {
		 return res
	 }
	 res = make([]string, 0)
	 makeStr("", root)
	 return res
 }
 
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 import "strconv"
 var res []string
 func binaryTreePaths(root *TreeNode) []string {
	 res = make([]string, 0)
	 makeStr("", root)
	 return res
 }
 
 func makeStr(str string, root *TreeNode) {
	 if root == nil {
		 return
	 }
	 if root.Left == nil && root.Right == nil {
		 res = append(res, str + strconv.Itoa(root.Val))
	 } else {
		 makeStr(str + strconv.Itoa(root.Val) + "->", root.Right)
		 makeStr(str + strconv.Itoa(root.Val) + "->", root.Left)
	 }
 }