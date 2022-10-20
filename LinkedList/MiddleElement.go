// Time Complexity : O(n) Space complexity O(1)
// Runtime : 3 ms	Memory : 2 MB

func middleNode(head *ListNode) *ListNode {
	var s *ListNode = head
	var f *ListNode = head


	for f!=nil && f.Next!=nil{
		s=s.Next
		f=f.Next.Next
	}
	return s

}
