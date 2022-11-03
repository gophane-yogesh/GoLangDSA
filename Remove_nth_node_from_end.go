// Time Complexity : O(n) two pass
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    curr := head
    count:=0
    
    for curr != nil{
        count++
        curr = curr.Next
        
    }
    
    if count - n == 0 {
        return head.Next
    }
    curr = head
    for i :=1; i<count-n; i++{
        curr = curr.Next
}
    curr.Next = curr.Next.Next
    
    
   
    return head
    
}
