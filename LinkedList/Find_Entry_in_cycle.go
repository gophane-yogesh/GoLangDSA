// Time Complexity : O(n) Space Complexity : O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    slow := head
    fast := head
    result :=head
    
    if head==nil || head.Next == nil {
        return nil
    }
    
    for  fast.Next!=nil && fast.Next.Next!=nil{
        slow=slow.Next
        fast=fast.Next.Next
        
        if slow == fast{
            for slow != result{
                slow=slow.Next
               result=result.Next
            }
            return result

        }
        
}   
    return nil
    
}
