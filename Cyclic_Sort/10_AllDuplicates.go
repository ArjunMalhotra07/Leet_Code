package Cyclic_Sort

import "fmt"

func FindDuplicatesHelper() {
	f := fmt.Println
	f("PROGRAM 4 : Find duplicates in an unsorted array ")
	var nums = []int{4, 8, 3, 2, 7, 8, 2, 3, 1}
	f("test array -- ", nums)
	ans := findDuplicates(nums)
	f("sorted", nums)
	f("duplicates", ans)
	f()
	f()
}
func findDuplicates(nums []int) []int {
	arr := []int{}
	i := 0
	for i < len(nums) {
		correct := nums[i] - 1
		if nums[i] != nums[correct] {
			swapThisFunc(nums, i, correct)

		} else {
			i++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			arr = append(arr, nums[i])

		}

	}
	return arr

}
func swapThisFunc(nums []int, i int, correct int) {

	temp := nums[i]
	nums[i] = nums[correct]
	nums[correct] = temp
}
