// Time : 6 ms	Space : 2.5 MB

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var curr *ListNode = head
    var prev *ListNode= nil
    
    for curr!=nil {
        var temp *ListNode = curr.Next
        curr.Next = prev
        prev = curr
        curr = temp
        
    }
    return prev
    
}
