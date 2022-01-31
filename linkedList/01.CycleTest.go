/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
       if head==nil{
        return false
    }
    q:=head
    p:=head
    for p!=nil && q!=nil && q.Next!=nil {
        p=p.Next
        q=q.Next.Next
        if p == q{
            return true
        }
    }
    return false
}
