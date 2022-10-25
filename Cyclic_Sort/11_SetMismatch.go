// https://leetcode.com/problems/set-mismatch/
package Cyclic_Sort

import "fmt"

func FindErrorNumsHelper() {
	f := fmt.Println
	f("PROGRAM 5 : To find duplicated element and finding missing element")
	arr := []int{2, 1, 5, 4, 5}
	f("test array", arr)
	ans := findErrorNums(arr)
	f(ans[0], "is duplicate")
	f(ans[1], "is missing element")
	f()
	f()
}
func findErrorNums(nums []int) []int {
	arr := []int{}

	i := 0
	for i < len(nums) {
		correct := nums[i] - 1
		if nums[i] != nums[correct] {
			swapMethod(nums, i, correct)
		} else {
			i++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			arr = append(arr, nums[i], i+1)
		}
	}

	return arr
}
func swapMethod(nums []int, i int, correct int) {
	temp := nums[i]
	nums[i] = nums[correct]
	nums[correct] = temp
}
