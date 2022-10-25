package main

import (
	"fmt"
)

func isPalindrome() {
	f := fmt.Println
	var x int
	f("PROGRAM 1 : Enter Number to check if Palindrome or Not")
	fmt.Scanln(&x)
	t := isPalindromeHelper(x)
	f(t)
	f()
	f()

}

func isPalindromeHelper(x int) bool {
	input := x
	var r int
	res := 0
	for x > 0 {
		r = x % 10
		res = res*10 + r
		x /= 10

	}
	if res == input {
		return true
	} else {
		return false
	}
}
