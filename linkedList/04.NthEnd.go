/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil { //or l.length==0
		return nil
	}
	currentNode := head
	total := 0
	for currentNode != nil {
		currentNode = currentNode.Next
		total++
	}
	if total == 1 {
		return nil
	}
	k := total - n + 1
	count := 1

	if k == 1 {
		head = head.Next
		return head
	}
	previousToDelete := head
	for count+1 != k {

		previousToDelete = previousToDelete.Next
		count += 1
	}
	previousToDelete.Next = previousToDelete.Next.Next
	return head
}