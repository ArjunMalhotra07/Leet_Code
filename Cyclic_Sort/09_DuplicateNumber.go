package Cyclic_Sort

import "fmt"

func SelectionSortFuncHelper() {
	f := fmt.Println
	f("PROGRAM 3 : To find missing element")
	var nums = []int{1, 2, 7, 4, 5, 6, 7, 8, 9, 10}
	ans, ans1 := SelectionSortFunc(nums)
	f(nums)
	if ans != ans1 {
		f(ans, " is extra in the array. ", ans1, " is missing")
	} else {
		f("Everything is in place")
	}
	f()
	f()
}
func SelectionSortFunc(nums []int) (int, int) {

	i := 0
	for i < len(nums) {
		correct := nums[i] - 1
		if nums[i] < len(nums) && nums[i] != nums[correct] {
			swapFunction(nums, i, correct)

		} else {
			i++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return nums[i], i + 1
		}

	}
	return len(nums), len(nums)

}
func swapFunction(nums []int, i int, correct int) {

	temp := nums[correct]
	nums[correct] = nums[i]
	nums[i] = temp
}
