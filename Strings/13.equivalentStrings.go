// https://leetcode.com/problems/check-if-two-string-arrays-are-equivalent/
package Strings

import "fmt"

func ArrayStringsAreEqualHelper() {
	f := fmt.Println
	f("PROGRAM 13 : Check If Two String Arrays are Equivalent")
	word1 := [3]string{"abc", "def", "g"}
	word2 := [1]string{"abcdefg"}
	f("Word 1", word1)
	f("Word 2", word2)
	f(arrayStringsAreEqual(word1[:], word2[:]))
	f()
	f()

}
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	var s1 string
	var s2 string

	for i := 0; i < len(word1); i++ {
		s1 += word1[i]
	}
	for i := 0; i < len(word2); i++ {
		s2 += word2[i]
	}
	return s1 == s2

}
