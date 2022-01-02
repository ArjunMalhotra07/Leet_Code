package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hlo")
	var n int
	fmt.Println("Enter Array Length")
	fmt.Scanln(&n)
	var n1 int
	fmt.Println("Enter Target")
	fmt.Scanln(&n1)
	fmt.Println("Enter Array")
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	t := twoSum(arr, n1)
	fmt.Printf("%v", t)

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
