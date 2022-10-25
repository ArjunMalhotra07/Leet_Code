//https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/submissions/
package Binary_Search

import "fmt"

func SearchRangeHelper() {
	f := fmt.Println
	f("PROBLEM 2 : Search last and first position for an element in sorted array")
	testArray := []int{-50, -25, 0, 4, 4, 100, 1000}
	target := 4
	f("Array -- ", testArray)
	f("Target -- ", target)
	f(SearchRange(testArray, 4))

	f()
	f()
}
func SearchRange(nums []int, target int) []int {
	start := search(nums, target, true)
	end := search(nums, target, false)

	return []int{start, end}

}

func search(nums []int, target int, findStartIndex bool) int {
	ans := -1
	start := 0
	end := len(nums) - 1

	for start <= end {

		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			ans = mid
			if findStartIndex {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
	}
	return ans
}
