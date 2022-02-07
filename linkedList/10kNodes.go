
func reverse(s *ListNode, e *ListNode) {
	current := s
	var prev, next *ListNode
	for prev != e {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	s := head
	e := head
	inc := k - 1
	for inc != 0 {
		e = e.Next
		if e == nil {
			return head
		}
		inc--
	}
	var nextHead *ListNode = reverseKGroup(e.Next, k)
	reverse(s, e)
	s.Next = nextHead
	return e
}
