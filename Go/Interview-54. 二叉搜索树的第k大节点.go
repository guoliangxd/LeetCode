package leetcode

/*给定一棵二叉搜索树，请找出其中第k大的节点。*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargest(root *TreeNode, k int) int {
	nums := make([]int, 0, k)
	getNums(root, &nums)
	return nums[k-1]
}

func getNums(root *TreeNode, nums *[]int) {
	if root.Right != nil {
		getNums(root.Right, nums)
	}
	if root != nil {
		*nums = append(*nums, root.Val)
	}
	if root.Left != nil {
		getNums(root.Left, nums)
	}
}
