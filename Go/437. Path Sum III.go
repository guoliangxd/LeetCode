package leetcode

/*You are given a binary tree in which each node contains an integer value.

Find the number of paths that sum to a given value.

The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).

The tree has no more than 1,000 nodes and the values are in the range -1,000,000 to 1,000,000.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/path-sum-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int
var target int

func pathSum(root *TreeNode, sum int) int {
	res = 0
	target = sum
	addPath(root, []int{})

	return res
}

func addPath(root *TreeNode, path []int) {
	if root == nil {
		return
	}

	path = append(path, root.Val)
	sum := 0
	for i := len(path) - 1; i >= 0; i-- {
		sum += path[i]
		if sum == target {
			res++
		}
	}

	addPath(root.Left, path)
	addPath(root.Right, path)
}
