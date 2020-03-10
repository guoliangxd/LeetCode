package leetcode

/*Consider all the leaves of a binary tree.  From left to right order, the values of those leaves form a leaf value sequence.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/leaf-similar-trees
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    list1 := make([]int, 0)
    list2 := make([]int, 0)
    getLeaf(&list1, root1)
    getLeaf(&list2, root2)
    if len(list1) != len(list2) {
        return false
    }

    for i := 0; i < len(list1); i++ {
        if list1[i] != list2[i] {
            return false
        }
    }
    return true
}

func getLeaf(list *[]int, root *TreeNode) {
    if root == nil {
        return 
    }

    if root.Left == nil && root.Right == nil {
        *list = append(*list, root.Val)
    } else {
        getLeaf(list, root.Left)
        getLeaf(list, root.Right)
    }
}