//https://leetcode.com/problems/find-first-palindromic-string-in-the-array/submissions/
package main

import (
	"fmt"
)

func main() {

	fmt.Println("hlo")
	words := []string{"abc", "car", "ada", "racecar", "cool"}
	//firstPalindrome(words)
	ans := firstPalindrome(words)
	fmt.Println(ans)
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
	var temp byte

	for i, j := 0, len(test_byte)-1; i < j; i, j = i+1, j-1 {
		temp = test_byte[i]
		test_byte[i] = test_byte[j]
		test_byte[j] = temp
	}
	if string(test_byte) == test {
		return true
	} else {
		return false
	}
}
