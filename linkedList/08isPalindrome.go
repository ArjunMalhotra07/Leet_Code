/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//Method 1
func isPalindrome(head *ListNode) bool {
	//Copy
	value := head
	mylist := list.New()
	for value != nil {
		mylist.PushBack(value.Val)
		value = value.Next
	}
	for element := mylist.Front(); element != nil; {
		fmt.Println(element.Value)
		element = element.Next()
	}
	//Reverse
	currentNode := head
	var prevNode *ListNode = nil
	var nextNode *ListNode = nil
	for currentNode != nil {
		nextNode = currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}
	head = prevNode
	//Check
	element := mylist.Front()
	testNode := head
	ans := true
	for element != nil && testNode != nil {
		if testNode.Val != element.Value {
			ans = false
		}
		element = element.Next()
		testNode = testNode.Next
	}
	return ans

}

