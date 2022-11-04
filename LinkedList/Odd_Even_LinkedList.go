// Time Complexity : O(n) Space Copmplexity : O(1)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    odd := head;
    even := head.Next
    evenhead := even
    
    for even!=nil && even.Next != nil {
        odd.Next = even.Next
        odd = odd.Next
        
        even.Next = odd.Next
        even = even.Next
    }
    odd.Next = evenhead
    
    return head
    
    
}
