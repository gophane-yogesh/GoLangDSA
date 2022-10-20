// Time Complexity : O(kN) where k is the number of linked list 
// Space Complexity : O(1)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    var dummy *ListNode = &ListNode{-999999,nil}
    ptr := dummy
    
    for _,item := range lists{
        ptr = mergeTwoLists(ptr,item)
        
        
    }
    return ptr.Next
    
}
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1==nil {
        return list2

    }
    if list2==nil{
        return list1
    }
    if list1.Val <= list2.Val{
        list1.Next=mergeTwoLists(list1.Next,list2)
        return list1
    }else{
        list2.Next=mergeTwoLists(list1,list2.Next)
        return list2
    }
    
    
    
}
