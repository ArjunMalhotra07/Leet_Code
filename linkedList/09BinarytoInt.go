
// type ListNode struct {
//     Val int
//     Next *ListNode
// }
func getDecimalValue(head *ListNode) int {
	currentNode := head
	length := 0
	for currentNode != nil {
		currentNode = currentNode.Next
		length += 1
	}
	sum := 0
	current := head
	for current != nil && length >= 0 {
		sum += int(math.Pow(float64(2), float64(length-1))) * (current.Val)
		length -= 1
		current = current.Next
	}
	return sum

}




