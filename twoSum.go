package main

import (
	"fmt"
)

func twoSumHelper() {
	f := fmt.Println
	var n int
	f("PROGRAM 5 : To find element pairs in array that equals a certain number")
	f("Enter Array Length")
	fmt.Scanln(&n)
	var n1 int
	f("Enter Target Sum that elements of your array should sum to")
	fmt.Scanln(&n1)
	f("Enter Array")
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	t := twoSum(arr, n1)
	f("Pairs at these position sum to the number you provided earlier")
	fmt.Printf("%v", t)
	f()
	f()
}

func twoSum(nums []int, target int) []int {
	var a []int
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				a = append(a, i, j)
			}
		}
	}
	return a
}
