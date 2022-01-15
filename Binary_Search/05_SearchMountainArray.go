//https://leetcode.com/problems/find-in-mountain-array/

package main

import "fmt"

func main() {

	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5, 1}
	res := peakIndexInMountainArray(a)
	fmt.Println(res)

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

	ans1 := binarySearchAsc(arr[0:start], 5)
	if ans1 != -1 {
		return ans1
	}
	ans2 := binarySearchDesc(arr[start:len(arr)-1], 5)
	return ans2
}
func binarySearchAsc(arr []int, target int) int {
	ans := -1
	start := 0
	end := len(arr) - 1
	for start <= end {
		//	mid:=(start+end)/2   Might be possible k ye exceeed kr jaye int ki limit
		mid := start + (end-start)/2

		if arr[mid] > target {
			end = mid - 1
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			ans = mid
		}

	}
	return ans
}

func binarySearchDesc(arr []int, target int) int {
	ans := -1
	start := 0
	end := len(arr) - 1
	for start <= end {
		//	mid:=(start+end)/2   Might be possible k ye exceeed kr jaye int ki limit
		mid := start + (end-start)/2

		if arr[mid] > target {
			start = mid + 1
		} else if arr[mid] < target {
			end = mid - 1
		} else {
			ans = mid
		}

	}
	return ans
}
