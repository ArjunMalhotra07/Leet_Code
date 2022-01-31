//Delete Middle of the Linked List
func deleteMiddle(head *ListNode) *ListNode {
	currentNode := head
	length := 0
	for currentNode != nil {
		length++
		currentNode = currentNode.Next
	}
	if length == 1 {
		return nil
	}
	mid := length / 2
	previousNode := head
	temp := 0
	for previousNode != nil {
		temp += 1
		if temp == mid {
			previousNode.Next = previousNode.Next.Next
			return head
		}
		previousNode = previousNode.Next
	}
	return nil
}

//OR.
func deleteMiddle(head *ListNode) *ListNode {

	slow := head
	fast := head
	if head.Next == nil {
		return nil
	}
	for fast != nil || fast.Next != nil {
		fast = fast.Next.Next
		if fast == nil || fast.Next == nil {
			slow.Next = slow.Next.Next
			return head
		}
		slow = slow.Next
	}

	return nil
}