package leetcode

/*Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in
 which the depth of the two subtrees of every node never differ by more than 1.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	var root *TreeNode
	if len(nums) == 0 {
		root = nil
	} else if len(nums) == 1 {
		root = &TreeNode{nums[0], nil, nil}
	} else {
		mid := len(nums) / 2
		root = &TreeNode{nums[mid], sortedArrayToBST(nums[:mid]), sortedArrayToBST(nums[mid+1:])}
	}
	return root
}
