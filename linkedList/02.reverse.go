func (l *linkedList) reverseLinkedList() *node {
        var previousNode, nextNode *node
        currentNode := l.head
        if currentNode == nil {
                return currentNode
        }
        for currentNode != nil {
                nextNode, currentNode.next = currentNode.next, previousNode

                previousNode, currentNode = currentNode, nextNode

        }
        l.head = previousNode
        return l.head
}
