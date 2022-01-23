//https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/
package main

import (
	"fmt"
	"strings"
)

func main() {

	p := fmt.Println
	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	p(mostWordsFound(sentences))  //Counting using Fields
	p(mostWordsFound1(sentences)) //Counting using Spaces

}
func mostWordsFound(sentences []string) int {

	length := len(sentences)
	var max int = 0

	for i := 0; i < length; i++ {
		str := sentences[i]

		words := strings.Fields(str)
		x := len(words)
		if x > max {
			max = x
		}

	}
	return max
}

// OR
func mostWordsFound1(sentences []string) int {

	length := len(sentences)
	var max int = 0

	for i := 0; i < length; i++ {
		str := sentences[i]
		x := 0

		for i := 0; i < len(strings.Trim(str, " ")); i++ {
			if string(str[i]) == " " {
				x++
			}
		}

		if x > max {
			max = x
		}

	}
	return max + 1
}
