//https://leetcode.com/problems/find-first-palindromic-string-in-the-array/submissions/
package Strings

import (
	"fmt"
	"strings"
)

func FirstPalindromeHelper() {
	p := fmt.Println
	p("PROGRAM 3 : Find first palindromic word in a string ")
	x := "cooc car ada racecar cool"
	p("String --  ", x)
	words := strings.Fields(x)
	p("First palindromic word ", firstPalindrome(words))
	p()
	p()
}

func firstPalindrome(words []string) string {
	for _, ans := range words {
		res := check(ans)
		if res {
			return ans
		}
	}
	return ""
}

func check(ans string) bool {
	test := ans
	test_byte := []byte(test)

	//1 by 1 krlia check
	for i, j := 0, len(test_byte)-1; i < j; i, j = i+1, j-1 {
		s1 := string(test_byte[i])
		s2 := string(test_byte[j])

		if s1 == s2 {
			return true
		} else {
			return false
		}

	}
	return false
}
