package leetcode

/*Given a binary tree, each node has value 0 or 1.  Each root-to-leaf path represents a binary number starting with the most significant bit.  For example, if the path is 0 -> 1 -> 1 -> 0 -> 1, then this could represent 01101 in binary, which is 13.

For all leaves in the tree, consider the numbers represented by the path from the root to that leaf.

Return the sum of these numbers.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int = 0

func sumRootToLeaf(root *TreeNode) int {
	res = 0
	sum(root, 0)
	return res
}

func sum(root *TreeNode, bin int) {
	if root == nil {
		return
	}

	bin = bin << 1
	if root.Left == nil && root.Right == nil {
		res += (bin + root.Val)
	} else {
		sum(root.Left, bin+root.Val)
		sum(root.Right, bin+root.Val)
	}
}
