package main

import "fmt"

func main() {
	var arr = []int{10, 20, 30, 50, 60, 70, 900, 1000, 5654, 9989, 298898}
	//n := ans(arr[0:], 30)
	//fmt.Println(ans(arr, 30))
	ans(arr, 30)
}
func ans(arr []int, target int) {
	start := 0
	end := 1
	//condition for target lie in the range

	for target > arr[end] {
		newstart := end + 1
		//Double box range
		//end= previous end + 2*size of box
		end = end + (end-start+1)*2
		start = newstart

	}
	result := binarySearch(arr, target, start, end)
	fmt.Println(result)
	//return result
}

func binarySearch(arr []int, target int, start int, end int) int {
	mid := start + (end-start)/2
	for start <= end {
		if target < arr[mid] {
			end = mid - 1
		} else if target > arr[mid] {
			start = mid + 1
		} else {
			return mid
		}

	}
	return -1
}
