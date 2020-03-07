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
    public int maxDepth(Node root) {
        if(root == null) {
            return 0;
        }
        int rlt = 0;
        for(int i = 0; i < root.children.size(); i++) {
            int len = maxDepth(root.children.get(i));
            if(len > rlt) {
                rlt = len;
            }
        }
        return rlt + 1;
    }
}