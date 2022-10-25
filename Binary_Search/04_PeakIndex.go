// https://leetcode.com/problems/peak-index-in-a-mountain-array/
package Binary_Search

import (
	"fmt"
)

func PeakIndexInMountainArrayHelper() {
	var n int
	f := fmt.Println
	f("PROBLEM 3 : Get index of the peak element in a mountain array")
	f("Enter Array Length")
	fmt.Scanln(&n)

	f("Enter Array such as [100,200,5,4,1]")
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &arr[i])
	}
	f(arr)

	ans := peakIndexInMountainArray(arr)
	f("Peak comes at ", ans)
	f()
	f()
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
