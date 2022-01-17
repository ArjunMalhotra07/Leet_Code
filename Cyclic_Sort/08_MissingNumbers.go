//Q268..    https://leetcode.com/problems/missing-number/

package main

import "fmt"

func main() {
	var nums = []int{4, 3, 2, 7, 8, 2, 3, 1}
	ans := findDisappearedNumbers(nums)
	fmt.Println(nums)
	fmt.Println(ans)
}

func findDisappearedNumbers(nums []int) []int {
	arr := []int{}
	i := 0
	for i < len(nums) {
		correct := nums[i] - 1
		if nums[i] != nums[correct] {
			swap(nums, i, correct)

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
func swap(nums []int, i int, correct int) {

	temp := nums[correct]
	nums[correct] = nums[i]
	nums[i] = temp
}
