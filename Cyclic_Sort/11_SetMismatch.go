package Cyclic_Sort

import "fmt"

func FindErrorNumsHelper() {
	arr := []int{1, 1}
	ans := findErrorNums(arr)
	fmt.Println(ans)
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
