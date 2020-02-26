package leetcode

/*Given the root node of a binary search tree,
return the sum of values of all nodes with value between L and R (inclusive).
The binary search tree is guaranteed to have unique values.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/range-sum-of-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	if root.Val >= L && root.Val <= R {
		return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
	} else if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	} else {
		return rangeSumBST(root.Left, L, R)
	}
}
