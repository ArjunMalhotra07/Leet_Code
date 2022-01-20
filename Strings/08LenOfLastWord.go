package main

import (
	"fmt"
	"strings"
)

func main() {
	x := fmt.Println

	s := "hi i love u "
	x(lengthOfLastWord(s))

}

func lengthOfLastWord(s string) int {
	x := fmt.Println
	words := strings.Fields(s)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	x(strings.Join(words, " "))
	return len(words[0])
}
