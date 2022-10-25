package Cyclic_Sort

import "fmt"

func SelectionSortFuncHelper() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7, 7}
	ans := SelectionSortFunc(nums)
	fmt.Println(nums)
	fmt.Println(ans)
}
func SelectionSortFunc(nums []int) int {

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

			return nums[i]
		}

	}
	return len(nums)

}
func swapFunction(nums []int, i int, correct int) {

	temp := nums[correct]
	nums[correct] = nums[i]
	nums[i] = temp
}
