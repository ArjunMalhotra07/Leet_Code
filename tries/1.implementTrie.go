type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}
func (this *Trie) Insert(word string) {

	currentNode := this
	for _, ch := range word {
		charIndex := ch - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Trie{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (this *Trie) Search(word string) bool {

	currentNode := this

	for _, ch := range word {
		charIndex := ch - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return currentNode.isEnd
}

func (this *Trie) StartsWith(word string) bool {

	currentNode := this
	for _, ch := range word {
		charIndex := ch - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	return true
}
