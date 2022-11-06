// Time Complexity : O(n) Space Complexity : O(1)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return helper(root) != -1
    
    
}

func helper(node *TreeNode) int {
    if node == nil {
        return 0
    }
    left := helper(node.Left)
    right := helper(node.Right)
    
    if left == -1 || right == -1 || abs(left-right) > 1 {
        return -1
    }
    
    return max(left,right) +1
    
}

func max(x, y int) int{
    if x > y{
        return x
    }
    return y
}

func abs(x int) int {
    if x < 0{
        return x *-1
    }
    return x
}
