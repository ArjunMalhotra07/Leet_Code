/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    currentNode:=head
    length:=0
    for currentNode!=nil{
        length++
        currentNode=currentNode.Next
    }
    if length==1{
        return nil
    }
    mid:= length/2
    previousNode:=head
    temp:=0
    for  previousNode!=nil{
        temp+=1
        if temp==mid{
      previousNode.Next=previousNode.Next.Next
      return head
        }
       previousNode=previousNode.Next
    }
    return nil
}
