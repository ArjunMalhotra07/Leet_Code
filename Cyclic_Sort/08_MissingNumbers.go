//Q268..    https://leetcode.com/problems/missing-number/

package Cyclic_Sort

import "fmt"

func FindDisappearedNumbersHelper() {

	f := fmt.Println
	f("PROGRAM 2 : To check for disappeard number in an array")
	var nums = []int{4, 3, 2, 7, 8, 2, 3, 1}
	f("Test Array -- ", nums)
	ans := findDisappearedNumbers(nums)
	f("Disappeard Numbers --- ", ans)
	f()
	f()
}
func findDisappearedNumbers(nums []int) []int {
	arr := []int{}
	i := 0
	for i < len(nums) {
		correct := nums[i] - 1
		if nums[i] != nums[correct] {
			swapFunct(nums, i, correct)

		} else {
			i++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			arr = append(arr, i+1)
			//arr[i] = i + 1
			//return i
		}

	}
	return arr
}
func swapFunct(nums []int, i int, correct int) {

	temp := nums[correct]
	nums[correct] = nums[i]
	nums[i] = temp
}
