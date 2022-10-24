// Time Complexity : O(n)
// Space Complexity : O(1)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    
    for root != nil{
        if p.Val > root.Val && q.Val > root.Val {
            root= root.Right
        }else if p.Val < root.Val && q.Val < root.Val {
            root= root.Left
        }else {
            return root
        }
    }
	return root
}
