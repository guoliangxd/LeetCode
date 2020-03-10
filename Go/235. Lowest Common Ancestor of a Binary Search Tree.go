package leetcode

/*Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	listP := make([]*TreeNode, 0)
	listQ := make([]*TreeNode, 0)
	TreeSearch(&listP, root, p)
	TreeSearch(&listQ, root, q)
	for i := 0; i < len(listP); i++ {
		for j := 0; j < len(listQ); j++ {
			if listP[i].Val == listQ[j].Val {
			   return listP[i]
			}
		}
	}
	return root
}

func TreeSearch(list *[]*TreeNode, root *TreeNode, target *TreeNode) {
   if root == nil {
	   return
   }

   if root.Val > target.Val {
	   TreeSearch(list, root.Left, target)
   } else if root.Val < target.Val {
	   TreeSearch(list, root.Right, target)
   }

   *list = append(*list, root)
}