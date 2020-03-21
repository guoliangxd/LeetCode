package leetcode

/*Given a Binary Search Tree and a target number, return true if there exist two elements in the BST such that their sum is equal to the given target.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var arr []int

func findTarget(root *TreeNode, k int) bool {
	arr = []int{}

	inOrder(root)

	for i := 0; i < len(arr)-1; i++ {
		target := k - arr[i]
		if binarySearch(i+1, target) {
			return true
		}
	}
	return false
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}

	inOrder(root.Left)
	arr = append(arr, root.Val)
	inOrder(root.Right)
}

func binarySearch(begin int, target int) bool {
	end := len(arr) - 1
	for begin <= end {
		mid := begin + (end-begin)/2
		if arr[mid] == target {
			return true
		} else if arr[mid] > target {
			end = mid - 1
		} else {
			begin = mid + 1
		}
	}
	return false
}

/*
var isFind bool
var treeRoot *TreeNode
func findTarget(root *TreeNode, k int) bool {
    isFind = false
    treeRoot = root
    inOrder(root, k)
    return isFind
}

func find(root *TreeNode, target int) bool {
    if root == nil {
        return false
    }

    res := false
    if target > root.Val {
        res = find(root.Right, target)
    } else if target == root.Val {
        res = true
    } else {
        res = find(root.Left, target)
    }
    return res
}

func inOrder(root *TreeNode, target int)  {
    if root == nil || isFind {
        return
    }

    inOrder(root.Left, target)
    if target - root.Val != root.Val {
        if find(treeRoot,target - root.Val) {
            isFind = true
        }
    }

    inOrder(root.Right, target)
}*/
