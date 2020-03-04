package leetcode

/*从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	rlt := make([][]int, 0, 0)
	PreOrderTree(1, &rlt, root)
	return rlt
}

func PreOrderTree(level int, arr *[][]int, root *TreeNode) {
	if root == nil {
		return
	}
	if len(*arr) < level {
		temp := make([]int, 0, 1)
		temp = append(temp, root.Val)
		*arr = append(*arr, temp)
	} else {
		(*arr)[level-1] = append((*arr)[level-1], root.Val)
	}
	PreOrderTree(level+1, arr, root.Left)
	PreOrderTree(level+1, arr, root.Right)
}
