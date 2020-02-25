package leetcode

/*给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。*/

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
        root = &TreeNode{nums[mid], sortedArrayToBST(nums[ : mid]), sortedArrayToBST(nums[mid + 1 : ])}
    }
    return root
}