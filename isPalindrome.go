package main

import (
	"fmt"
)

func isPalindrome() {
	var x int
	fmt.Println("Enter Number")
	fmt.Scanln(&x)
	t := isPalindromeHelper(x)
	fmt.Println(t)
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
