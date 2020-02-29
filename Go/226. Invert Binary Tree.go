package leetcode

/*Invert a binary tree.*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func invertTree(root *TreeNode) *TreeNode {
    if root != nil {
        temp := root.Left
        root.Left = root.Right
        root.Right = temp
        invertTree(root.Left)
        invertTree(root.Right)
    }
    return root
}s
