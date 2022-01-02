package main

import (
	"fmt"
)

func main() {

	var x int
	fmt.Println("Enter Number")
	fmt.Scanln(&x)
	isPalindrome(x)
	t := isPalindrome(x)
	fmt.Println(t)
}
func isPalindrome(x int) bool {
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
