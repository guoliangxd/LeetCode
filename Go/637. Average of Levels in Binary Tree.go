package leetcode

/*Given a non-empty binary tree, return the average value of the nodes on each level in the form of an array.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/average-of-levels-in-binary-tree
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

func averageOfLevels(root *TreeNode) []float64 {
	rlt = make([][]int, 0)
	treePreOrder(root, 1)

	res := make([]float64, len(rlt))
	for i := 0; i < len(rlt); i++ {
		sum := 0.0
		for j := 0; j < len(rlt[i]); j++ {
			sum += float64(rlt[i][j])
		}
		res[i] = sum / float64(len(rlt[i]))
	}
	return res
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
