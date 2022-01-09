// https://leetcode.com/problems/peak-index-in-a-mountain-array/

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