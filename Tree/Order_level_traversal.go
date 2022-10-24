// Time Complexity : O(n)
// Space Complexity : O(n)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
   if root == nil {
        return [][]int{}
    }
   var queue []*TreeNode
    var result [][]int
    queue=append(queue,root)
    
    for len(queue)>0{
        var res []int
        currlen := len(queue)
        
        for i:=0;i<currlen ;i++{
            element:=queue[0]
            
            queue=queue[1:]
            if element.Left!=nil{
                queue = append(queue,element.Left)
            }
            if element.Right!=nil{
                queue = append(queue,element.Right)
            }
            res = append(res,element.Val)
            
        }
        result = append(result,res)
        
    }
    return result
}


