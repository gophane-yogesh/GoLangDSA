
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Recurssive Solution
// Time Complexity : O(n)
func inorderTraversal(root *TreeNode) []int {
    var result []int
    
    helper(root, &result)
    return result
    
}

func helper(node *TreeNode, result *[]int){
    if node == nil{
        return
    }
   
    helper(node.Left,result)
    
    *result = append(*result, node.Val)
    helper(node.Right,result)
}
