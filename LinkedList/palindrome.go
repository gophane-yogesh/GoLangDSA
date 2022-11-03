// Time Complexity : O(n) Space Complexity : O(n)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    var st []*ListNode
    curr := head
    
    for curr!= nil{
        st = append(st,curr)
        curr = curr.Next
    }
    curr = head
    for len(st)>0{
        if st[len(st)-1].Val == curr.Val{
            st = st[:len(st)-1]
            curr = curr.Next
            
        }else{
            return false
        }
        

    }
    return true
    
}
