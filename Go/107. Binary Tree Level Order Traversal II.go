package leetcode

/*Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var rlt [][]int

func levelOrderBottom(root *TreeNode) [][]int {
	rlt = make([][]int, 0)
	treePreOrder(root, 1)

	for i := 0; i < len(rlt)/2; i++ {
		rlt[i], rlt[len(rlt)-i-1] = rlt[len(rlt)-i-1], rlt[i]
	}
	return rlt
}

func treePreOrder(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(rlt) < level {
		temp := make([]int, 0)
		rlt = append(rlt, temp)
	}

	rlt[level-1] = append(rlt[level-1], root.Val)
	treePreOrder(root.Left, level+1)
	treePreOrder(root.Right, level+1)
}
