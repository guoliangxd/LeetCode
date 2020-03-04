/*Given an n-ary tree, return the preorder traversal of its nodes' values.

Nary-Tree input serialization is represented in their level order traversal, 
each group of children is separated by the null value (See examples).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, List<Node> _children) {
        val = _val;
        children = _children;
    }
};
*/
class Solution {
    public List<Integer> preorder(Node root) {
        List<Integer> rlt = new ArrayList<>();
        preorder(root, rlt);
        return rlt;
    }

    private void preorder(Node root, List<Integer> list) {
        if(root == null) {
            return;
        }
        list.add(root.val);
        for(int i = 0; i < root.children.size(); i++) {
            preorder(root.children.get(i), list);
        }
    }
}