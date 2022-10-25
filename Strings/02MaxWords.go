//https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/
package Strings

import (
	"fmt"
	"strings"
)

func CountUsingFieldsHelper() {
	p := fmt.Println
	p("PROGAM 2 : maximum-number-of-words-found-in-sentences")
	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	p("Original array -- ")
	print(sentences)
	p("Counting using Fields ", countUsingFields(sentences))
	p("Counting using spaces ", countUsingSpaces(sentences))
	p()
	p()
}
func print(sentences []string) {
	p := fmt.Println
	for i := 0; i < len(sentences); i++ {
		fmt.Print("Element ", i+1, " -     ", sentences[i])
		p()
	}

}
func countUsingFields(sentences []string) int {

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
func countUsingSpaces(sentences []string) int {

	length := len(sentences)
	var max int = 0

	for i := 0; i < length; i++ {
		str := sentences[i]
		testString := strings.TrimSpace(str)
		x := 0

		for i := 0; i < len(testString); i++ {
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
