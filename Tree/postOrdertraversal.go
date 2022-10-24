// Time COmplexity : O(n)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    var result []int
    helper(root, &result)
    return result
    
}

func helper(node *TreeNode, result *[]int){
    if node != nil{
        helper(node.Left, result)
        helper(node.Right, result)
        *result =append(*result, node.Val)
    }
}
