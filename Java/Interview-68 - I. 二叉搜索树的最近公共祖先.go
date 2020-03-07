package leetcode

/*给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-sou-suo-shu-de-zui-jin-gong-gong-zu-xian-lcof
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
 class Solution {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        TreeNode rlt = null;
        List<TreeNode> listP = new ArrayList<>();
        List<TreeNode> listQ = new ArrayList<>();
        treeSearch(listP, root, p);
        treeSearch(listQ, root, q);

        for(int i = 0; i < listP.size(); i++) {
            if(listQ.contains(listP.get(i))) {
                rlt = listP.get(i);
                break;
            }
        }
        return rlt;
    }

    private void treeSearch(List<TreeNode> list, TreeNode root, TreeNode target) {
        if(root == null) {
            return;
        }

        if(root.val < target.val) {
            treeSearch(list, root.right, target);
        } else if(root.val > target.val){
            treeSearch(list, root.left, target);
        }

        list.add(root);
    }
}