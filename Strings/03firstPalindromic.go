//https://leetcode.com/problems/find-first-palindromic-string-in-the-array/submissions/
package main

import (
	"fmt"
	"strings"
)

func main() {

	p := fmt.Println
	x := "Hello my name is arjun"
	p(strings.Fields(x))

	words := []string{"cooc", "car", "ada", "racecar", "cool"}
	p(firstPalindrome(words))
	p(check(x))
	p(palindrome(x))

}

func firstPalindrome(words []string) string {
	for _, ans := range words {
		res := check(ans)
		if res == true {
			return ans
		}
	}
	return ""
}

// func check(ans string) bool {
// 	test := ans
// 	test_byte := []byte(test)
// 	var temp1 byte
// 	var temp2 byte
// 	//1 by 1 krlia check
// 	for i, j := 0, len(test_byte)-1; i < j; i, j = i+1, j-1 {
// 		temp1 = test_byte[i]
// 		temp2 = test_byte[j]

// 		if temp1 == temp2 {
// 			return true
// 		} else {
// 			return false
// 		}

// 	}
// 	return false
// }

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

func palindrome(ans string) string {
	test := ans
	test_byte := []byte(test)

	//1 by 1 krlia check
	for i, j := 0, len(test_byte)-1; i < j; i, j = i+1, j-1 {
		s1 := string(test_byte[i])
		s2 := string(test_byte[j])

		var temp string
		temp = s1
		s1 = s2
		s2 = temp

	}
	return string(test_byte)
}
