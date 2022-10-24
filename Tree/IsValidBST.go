// Time Complexity :O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return valid(root, math.MinInt, math.MaxInt)
    
}

func valid(root *TreeNode, left int, right int) bool{
    if root == nil {
        return true
    }
    if root.Val >= right || root.Val <=left {
        return false
    } 
    return valid(root.Left, left, root.Val) && valid (root.Right, root.Val, right)
}
