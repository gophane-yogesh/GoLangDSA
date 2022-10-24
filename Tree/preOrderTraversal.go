// Time Complexity : O(n)


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// #1 Solution
func preorderTraversal(root *TreeNode) []int {
   
    var result []int
    helper(root,&result)
    return result
      
}

func helper(root *TreeNode, result *[]int){
    if root == nil{
        return
    }
    *result = append(*result,root.Val)
    helper(root.Left,result)
    helper(root.Right,result)
    
}

// #2 Solution
func preorderTraversal(root *TreeNode) []int {
   
    var result []int
    if root != nil{
        result = append(result,root.Val)
        return append(append(result,preorderTraversal(root.Left)...),preorderTraversal(root.Right)...)
    }
    
    return result
      
}


// Iterative Solution
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
   
    var result []int
    
    var stack []*TreeNode
    stack = append(stack,root)
    for len(stack)>0 {
        l := len(stack)-1
        element := stack[l]
        stack = stack[:l]
        if element != nil{
            result = append(result,element.Val)
            stack = append(stack, element.Right)
            stack = append(stack, element.Left)

        }
    }
    
    return result
      
}

