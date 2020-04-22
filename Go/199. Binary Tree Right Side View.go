package leetcode

/*Given a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-right-side-view
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var depth int
var res []int

func rightSideView(root *TreeNode) []int {
	depth = 0
	res = make([]int, 0)
	preOrder(root, 0)
	return res
}

func preOrder(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == depth {
		res = append(res, root.Val)
		depth++
	}
	level++
	preOrder(root.Right, level)
	preOrder(root.Left, level)
}
