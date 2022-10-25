//https://leetcode.com/problems/find-numbers-with-even-number-of-digits/submissions
package main

import (
	"fmt"
	"strconv"
)

func findNumbersHelper() {
	f := fmt.Println
	var n int
	f("Enter Array Length")
	fmt.Scanln(&n)

	f("Enter Array to Find Numbers with Even Number of Digits")
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	ans := findNumbers(arr)
	f(ans)
	f()
	f()
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
