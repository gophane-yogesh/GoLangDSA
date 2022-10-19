// Time Complexity O(n) and space Complexity O(n)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    result := &ListNode{}
    var p *ListNode = result
    var carry int = 0
    
    
    for l1!=nil && l2!=nil{
        sum := l1.Val + l2.Val + carry
        if sum>=10{
            carry=1
            sum=sum%10
        }else{
            carry=0
        }
        p.Next = &ListNode{sum,nil}
        l1=l1.Next
        l2= l2.Next
        p = p.Next
       
}
    for l1!= nil{
        sum := l1.Val  + carry
        if sum>=10{
            carry=1
            sum=sum%10
        }else{
            carry=0
        }
        p.Next = &ListNode{sum,nil}
        l1=l1.Next
        p = p.Next
    }
     for l2!= nil{
        sum := l2.Val  + carry
        if sum>=10{
            carry=1
            sum=sum%10
        }else{
            carry=0
        }
        p.Next = &ListNode{sum,nil}
        l2=l2.Next
        p = p.Next
    }
    if carry == 1 {
        p.Next = &ListNode {carry,nil}
        p = p.Next
    }
    return result.Next
    
}
