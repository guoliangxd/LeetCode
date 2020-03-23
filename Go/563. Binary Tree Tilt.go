package leetcode

/*Given a binary tree, return the tilt of the whole tree.

The tilt of a tree node is defined as the absolute difference between the sum of all left subtree node values and the sum of all right subtree node values. Null node has tilt 0.

The tilt of the whole tree is defined as the sum of all nodes' tilt.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-tilt
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}

	_, res := sum(root)
	return res
}

func sum(root *TreeNode) (int, int) { //总和，差和
	if root == nil {
		return 0, 0
	}

	lSum, lDif := sum(root.Left)
	rSum, rDif := sum(root.Right)
	return lSum + rSum + root.Val, int(math.Abs(float64(lSum)-float64(rSum))) + lDif + rDif
}
