package leetcode

/*请完成一个函数，输入一个二叉树，该函数输出它的镜像。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
    } else {
        temp := root.Left
        root.Left = root.Right
        root.Right = temp
            
        mirrorTree(root.Left) 
        mirrorTree(root.Right)

    }
    return root
}