//Return middle Node
func middleNode(head *ListNode) *ListNode {
	currentNode := head
	length := 0
	for currentNode != nil {
		length++
		currentNode = currentNode.Next
	}
	if length == 1 {
		return head
	}
	mid := length / 2
	currentNode1 := head
	temp := 0
	for temp != mid {
		temp++
		currentNode1 = currentNode1.Next
	}
	return currentNode1
}

//Or
func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	if head.Next == nil {
		return head
	}
	for fast != nil || fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == nil || fast.Next == nil {
			return slow
		}

	}
	return nil
}