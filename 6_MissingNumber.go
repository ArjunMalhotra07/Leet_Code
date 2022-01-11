//Q268..    https://leetcode.com/problems/missing-number/

package main

import "fmt"

func main() {
	var nums = []int{3, 0, 1}
	ans := SelectionSort(nums)
	fmt.Println(nums)
	fmt.Println(ans)
}

func SelectionSort(nums []int) int {

	i := 0
	for i < len(nums) {
		correct := nums[i]
		if nums[i] < len(nums) && nums[i] != nums[correct] {
			swap(nums, i, correct)

		} else {
			i++
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {

			return i
		}

	}
	return len(nums)

}
func swap(nums []int, i int, correct int) {

	temp := nums[correct]
	nums[correct] = nums[i]
	nums[i] = temp
}
