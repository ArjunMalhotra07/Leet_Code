//https://leetcode.com/problems/find-first-palindromic-string-in-the-array/submissions/
package main

import (
	"fmt"
)

func main() {

	p := fmt.Println
	words := []string{"racecar", "car", "ada", "racecar", "cool"}
	p(firstPalindrome(words))

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
func check(ans string) bool {
	test := ans
	test_byte := []byte(test)
	var temp1 byte
	var temp2 byte
	//Ye krlia check 1 by 1 OR Reverse the word and then check
	for i, j := 0, len(test_byte)-1; i < j; i, j = i+1, j-1 {
		temp1 = test_byte[i]
		temp2 = test_byte[j]

		if temp1 == temp2 {
			return true
		}

	}
	return false
}
