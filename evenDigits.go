//https://leetcode.com/problems/find-numbers-with-even-number-of-digits/submissions
package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int
	fmt.Println("Enter Array Length")
	fmt.Scanln(&n)

	fmt.Println("Enter Array")
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	for i := 0; i < n; i++ {
		fmt.Println(arr[i])
	}

	ans := findNumbers(arr)
	fmt.Println(ans)

}

func findNumbers(arr []int) int {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			arr[i] *= -1
		}
		y := strconv.Itoa(arr[i])
		length := len(y)
		if length%2 == 0 {
			count++
		}
	}
	return count
}
