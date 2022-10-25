// https://leetcode.com/problems/peak-index-in-a-mountain-array/
package Binary_Search

import (
	"fmt"
)

func PeakIndexInMountainArrayHelper() {
	var n int
	fmt.Println("Enter Array Length")
	fmt.Scanln(&n)

	fmt.Println("Enter Array")
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	fmt.Println(arr)

	ans := peakIndexInMountainArray(arr)
	fmt.Println(ans)
}
func peakIndexInMountainArray(arr []int) int {
	start := 0
	end := len(arr) - 1

	for start < end {
		mid := start + (end-start)/2

		if arr[mid] > arr[mid+1] { //u r in dec part of array
			//this may be the ans but still look at left
			end = mid
		} else {
			start = mid + 1
		}

	}
	return start
}
