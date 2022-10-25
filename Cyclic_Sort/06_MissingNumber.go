//Q268..    https://leetcode.com/problems/missing-number/

package Cyclic_Sort

import "fmt"

func SelectionSortHelperFunc() {
	f := fmt.Println
	f("PROGRAM 1 : To check for the missing number in Array")
	var nums = []int{0, 2, 3, 7, 5, 6, 4, 7}
	f("Array -- ", nums)
	ans := SelectionSort(nums)
	f("Sorted Array -- ", nums)
	f("Missing -- ", ans)
	f()
	f()
}

func SelectionSort(nums []int) int {

	i := 0
	for i < len(nums) {
		correct := nums[i]
		if nums[i] < len(nums) && nums[i] != nums[correct] { //N se chota rhe hmara i
			swapFunc(nums, i, correct)

		} else {
			i++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
			//return nums[i]
		}

	}
	return len(nums)

}
func swapFunc(nums []int, i int, correct int) {

	temp := nums[correct]
	nums[correct] = nums[i]
	nums[i] = temp
}
